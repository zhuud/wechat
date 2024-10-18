package save

import (
	"context"
	"errors"
	"fmt"
	"rpc/internal/config"
	"strings"
	"time"

	"reflect"
	"rpc/internal/svc"
	"rpc/model"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"github.com/avast/retry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhuud/go-library/utils"
)

type SaveExternalUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	HistoryWechatModel model.TbPrivateDomainUserModel // 兼容历史数据 新项目不需要
}

func NewSaveExternalUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExternalUserLogic {
	return &SaveExternalUserLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		HistoryWechatModel: model.NewTbPrivateDomainUserModel(sqlx.NewMysql(svcCtx.Config.WechatDb.HistoryWechatDataSource)),
	}
}

func (s *SaveExternalUserLogic) Save(crop string, wechatExternalUser *response.ResponseExternalContact, mode string) error {
	if wechatExternalUser == nil || wechatExternalUser.ExternalContact == nil {
		return errors.New("externalUser is nil")
	}

	// 外部用户信息主表保存
	ts := time.Now().Local()
	externalUser := &model.TbExternalUser{
		ExternalUserid: wechatExternalUser.ExternalContact.ExternalUserID,
		Unionid:        wechatExternalUser.ExternalContact.UnionID,
		Type:           int64(wechatExternalUser.ExternalContact.Type),
		Name:           wechatExternalUser.ExternalContact.Name,
		Avatar:         wechatExternalUser.ExternalContact.Avatar,
		Gender:         int64(wechatExternalUser.ExternalContact.Gender),
		CorpName:       wechatExternalUser.ExternalContact.CorpName,
		CorpFullName:   wechatExternalUser.ExternalContact.CorpFullName,
		Position:       wechatExternalUser.ExternalContact.Position,
		Status:         model.TbExternalUserNormalStatus,
		UpdatedAt:      ts,
	}
	err := s.SaveExternalUser(externalUser)
	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户信息失败, error: %v", err))
	}

	// 外部用户信息附属信息保存
	// 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段
	if wechatExternalUser.ExternalContact.ExternalProfile != nil && externalUser.Type == 2 {
		externalUserAttrList := make([]*model.TbExternalUserAttribute, 0)
		for _, item := range wechatExternalUser.ExternalContact.ExternalProfile.ExternalAttr {
			ext := ""
			switch item.Type {
			case model.AttributeTypeText:
				ext, _ = jsoniter.MarshalToString(item.Text)
			case model.AttributeTypeWeb:
				ext, _ = jsoniter.MarshalToString(item.Web)
			case model.AttributeTypeMiniprogram:
				ext, _ = jsoniter.MarshalToString(item.MiniProgram)
			}

			externalUserAttrList = append(externalUserAttrList, &model.TbExternalUserAttribute{
				ExternalUserid: wechatExternalUser.ExternalContact.ExternalUserID,
				AttributeType:  int64(item.Type),
				AttributeValue: item.Name,
				Extension:      ext,
				Status:         model.TbExternalUserAttrNormalStatus,
				CreatedAt:      ts,
				UpdatedAt:      ts,
			})
		}
		err = s.SaveExternalUserAttr(externalUser.ExternalUserid, externalUserAttrList)
	}
	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户扩展信息失败, error: %v", err))
	}

	// 外部用户信息关联信息保存
	// 没有关注信息结束 正常数据不会存在
	if wechatExternalUser.FollowInfo == nil {
		return errors.New(fmt.Sprintf("外部用户关注人信息为空, externalUserid: %s", externalUser.ExternalUserid))
	}

	externalUserFollow := &model.TbExternalUserFollow{
		ExternalUserid:    wechatExternalUser.ExternalContact.ExternalUserID,
		Unionid:           wechatExternalUser.ExternalContact.UnionID,
		Userid:            wechatExternalUser.FollowInfo.UserID,
		Crop:              crop,
		OperUserid:        wechatExternalUser.FollowInfo.OperUserID,
		AddWay:            int64(wechatExternalUser.FollowInfo.AddWay),
		State:             wechatExternalUser.FollowInfo.State,
		StateChannel:      "",
		StateChannelValue: "",
		Remark:            wechatExternalUser.FollowInfo.Remark,
		// 手机号加密存储
		RemarkMobiles:    utils.XorEncrypt(strings.Join(wechatExternalUser.FollowInfo.RemarkMobiles, ","), s.svcCtx.Config.Secret.Cryptographic),
		Description:      wechatExternalUser.FollowInfo.Description,
		RemarkCorpName:   wechatExternalUser.FollowInfo.RemarkCorpName,
		RemarkPicMediaid: "", // TODO 还没有此字段
		Status:           model.TbExternalUserFollowNormalStatus,
		CreatedAt:        ts,
		UpdatedAt:        ts,
	}
	// 加c渠道自定义数据处理
	s.fmtJoinCCallbackData(externalUserFollow)
	// 兼容历史数据关系状态 后续下掉
	if mode == "cmd" {
		s.getFromHistory(externalUserFollow)
	}

	err = s.SaveExternalUserFollow(externalUserFollow)
	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户关系信息失败, error: %v", err))
	}

	// 外部用户关系属性信息
	externalUserFollowAttrList := make([]*model.TbExternalUserFollowAttribute, 0)
	for _, tag := range wechatExternalUser.FollowInfo.Tags {
		tagJson, _ := jsoniter.MarshalToString(tag)
		externalUserFollowAttrList = append(externalUserFollowAttrList, &model.TbExternalUserFollowAttribute{
			ExternalUserid: wechatExternalUser.ExternalContact.ExternalUserID,
			Userid:         wechatExternalUser.FollowInfo.UserID,
			Crop:           crop,
			AttributeType:  model.AttributeTypeRemarkTag,
			AttributeValue: tag.TagID,
			Extension:      tagJson,
			Status:         model.TbExternalUserFollowNormalStatus,
			CreatedAt:      ts,
			UpdatedAt:      ts,
		})
	}
	err = s.SaveExternalUserFollowAttr(externalUserFollow.ExternalUserid, externalUserFollow.Userid, externalUserFollow.Crop, externalUserFollowAttrList)
	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户关系信息失败, error: %v", err))
	}

	return nil
}

func (s *SaveExternalUserLogic) SaveExternalUser(externalUser *model.TbExternalUser) error {
	if externalUser == nil || len(externalUser.ExternalUserid) == 0 {
		return errors.New("SaveExternalUser is nil")
	}

	var (
		dbExternalUser *model.TbExternalUser
		err            error
	)

	// 外部用户信息主表保存
	_ = retry.Do(func() error {
		dbExternalUser, err = s.svcCtx.ModelExternalUser.FindOne(s.ctx, externalUser.ExternalUserid)
		return err
	}, retry.Attempts(3))

	if dbExternalUser != nil {
		externalUser.CreatedAt = dbExternalUser.CreatedAt
		dbExternalUser.UpdatedAt = externalUser.UpdatedAt
		// 判断数据是否一致 一致不更新了
		if !reflect.DeepEqual(externalUser, dbExternalUser) {
			err = s.svcCtx.ModelExternalUser.Update(s.ctx, externalUser)
		} else {
			s.Infof("SaveExternalUser 数据一致不更新, externalUser: %v", externalUser)
		}

	} else {
		_, err = s.svcCtx.ModelExternalUser.Insert(s.ctx, externalUser)
	}

	return err
}

func (s *SaveExternalUserLogic) SaveExternalUserAttr(externalUserid string, externalUserAttrList []*model.TbExternalUserAttribute) error {
	err := s.svcCtx.ModelExternalUserAttribute.DeleteByExternalUserId(s.ctx, externalUserid)
	if err != nil {
		return errors.New(fmt.Sprintf("SaveExternalUserAttr.DeleteByExternalUserId error: %s", err.Error()))
	}

	for _, item := range externalUserAttrList {
		_, err = s.svcCtx.ModelExternalUserAttribute.Insert(s.ctx, item)
		if err != nil {
			return errors.New(fmt.Sprintf("SaveExternalUserAttr.Insert externalUserid:%s, error: %s", externalUserid, err.Error()))
		}
	}

	return err
}

func (s *SaveExternalUserLogic) fmtJoinCCallbackData(externalUserFollow *model.TbExternalUserFollow) {
	stateList := make([]string, 0)
	// "state": "99_9999#1#01234567890123456789"  或者 "99_9999#0#{"hid":15731}"
	// 充分校验下 防止正常渠道里带#号
	if !strings.Contains(externalUserFollow.State, "#") {
		return
	}
	if stateList = strings.Split(externalUserFollow.State, "#"); len(stateList) == 3 && utils.ArrayIn(stateList[1], []string{"0", "1"}) {
		externalUserFollow.StateChannel = stateList[0]
		externalUserFollow.StateChannelValue = stateList[2]
		if stateList[1] == "1" {
			// TODO 从缓存里面取具体内容
		}
	}
}

func (s *SaveExternalUserLogic) SaveExternalUserFollow(externalUserFollow *model.TbExternalUserFollow) error {
	if externalUserFollow == nil || len(externalUserFollow.ExternalUserid) == 0 || len(externalUserFollow.Userid) == 0 {
		return errors.New("SaveExternalUserFollow is nil")
	}

	var (
		dbExternalUserFollow *model.TbExternalUserFollow
		err                  error
	)
	err = retry.Do(func() error {
		dbExternalUserFollow, err = s.svcCtx.ModelExternalUserFollow.FindOneByExternalUserIdAndUserId(s.ctx, externalUserFollow.ExternalUserid, externalUserFollow.Userid, externalUserFollow.Crop)
		return err
	}, retry.Attempts(3))
	if err != nil {
		return errors.New(fmt.Sprintf("获取旧外部用户关系信息失败, error: %v", err))
	}

	if dbExternalUserFollow != nil {
		externalUserFollow.Seq = dbExternalUserFollow.Seq
		externalUserFollow.CreatedAt = dbExternalUserFollow.CreatedAt
		err = s.svcCtx.ModelExternalUserFollow.Update(s.ctx, dbExternalUserFollow)
	} else {
		_, err = s.svcCtx.ModelExternalUserFollow.Insert(s.ctx, externalUserFollow)
	}

	return err
}

func (s *SaveExternalUserLogic) SaveExternalUserFollowAttr(externalUserid, userid, crop string, externalUserFollowAttrList []*model.TbExternalUserFollowAttribute) error {

	err := s.svcCtx.ModelExternalUserFollowAttribute.DeleteByExternalUserIdAndUserId(s.ctx, externalUserid, userid, crop)
	if err != nil {
		return errors.New(fmt.Sprintf("SaveExternalUserFollowAttr.DeleteByExternalUserIdAndUserId error: %s", err.Error()))
	}

	for _, item := range externalUserFollowAttrList {
		_, err = s.svcCtx.ModelExternalUserFollowAttribute.Insert(s.ctx, item)
		if err != nil {
			return errors.New(fmt.Sprintf("SaveExternalUserFollowAttr.Insert externalUserid:%s, userid:%s, crop:%s, error: %s", externalUserid, userid, crop, err.Error()))
		}
	}

	return nil
}

func (s *SaveExternalUserLogic) getFromHistory(externalUserFollow *model.TbExternalUserFollow) {
	// 1新氧医美  2新氧优享 3新氧咨询
	dbHistory, _ := s.HistoryWechatModel.FindOneByExternalUserIdAndUserId(s.ctx, externalUserFollow.ExternalUserid, externalUserFollow.Userid, config.MapCropToHistoryPlatform(externalUserFollow.Crop))
	if dbHistory != nil {
		externalUserFollow.ChatAgreeStatus = dbHistory.ChatAgreeStatus
		if dbHistory.LastChatTime.Format(time.DateTime) != "1000-01-01 00:00:00" && dbHistory.LastChatTime.Format(time.DateTime) != "0001-01-01 00:00:00" {
			externalUserFollow.LastChatTime = dbHistory.LastChatTime
		}
		if dbHistory.AddTime.Format(time.DateTime) != "1000-01-01 00:00:00" && dbHistory.AddTime.Format(time.DateTime) != "0001-01-01 00:00:00" {
			externalUserFollow.CreatedAt = dbHistory.AddTime
		}
		if dbHistory.Status != 1 && dbHistory.BlacklistTime.Format(time.DateTime) != "1000-01-01 00:00:00" && dbHistory.BlacklistTime.Format(time.DateTime) != "0001-01-01 00:00:00" {
			externalUserFollow.Status = model.TbExternalUserFollowExternalUserDeleteStatus
			if dbHistory.BlacklistType == "staff" {
				externalUserFollow.Status = model.TbExternalUserFollowStaffDeleteStatus
			}
			externalUserFollow.DeletedAt = dbHistory.BlacklistTime
		}
	}
}
