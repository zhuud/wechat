package save

import (
	"context"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/zhuud/go-library/utils"
	"rpc/internal/svc"
	"rpc/model"
	"strings"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveExternalUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExternalUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExternalUserLogic {
	return &SaveExternalUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *SaveExternalUserLogic) Save(crop string, externalUser *response.ResponseExternalContact) error {
	if externalUser == nil || externalUser.ExternalContact == nil {
		return errors.New("externalUser is nil")
	}

	var (
		dbExternalUser       *model.TbExternalUser
		dbExternalUserFollow *model.TbExternalUserFollow
		err                  error
	)

	// 外部用户信息主表保存
	_ = retry.Do(func() error {
		dbExternalUser, err = s.svcCtx.ModelExternalUser.FindOne(s.ctx, externalUser.ExternalContact.ExternalUserID)
		return err
	}, retry.Attempts(3))

	ts := time.Now().Local()
	externalUserData := &model.TbExternalUser{
		ExternalUserid: externalUser.ExternalContact.ExternalUserID,
		Unionid:        externalUser.ExternalContact.UnionID,
		Type:           uint64(externalUser.ExternalContact.Type),
		Name:           externalUser.ExternalContact.Name,
		Avatar:         externalUser.ExternalContact.Avatar,
		Gender:         uint64(externalUser.ExternalContact.Gender),
		CorpName:       externalUser.ExternalContact.CorpName,
		CorpFullName:   externalUser.ExternalContact.CorpFullName,
		Position:       externalUser.ExternalContact.Position,
		Status:         model.TbExternalUserNormalStatus,
		UpdatedAt:      ts,
	}
	if dbExternalUser != nil {
		// TODO 判断数据是否一致 一致不更新了
		err = s.svcCtx.ModelExternalUser.Update(s.ctx, externalUserData)
	} else {
		externalUserData.CreatedAt = ts
		_, err = s.svcCtx.ModelExternalUser.Insert(s.ctx, externalUserData)
	}
	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户信息失败, error: %v", err))
	}

	// 外部用户信息附属信息保存
	// 因为微信批量方法里返回的关注人是一个，会出现外部用户多次出现的情况 attr属性没法判断是不是要更新只能删除新写
	// 跑数据不用每次都删除新建
	if _, ok := s.svcCtx.LocalCache.Get("external_user_attr_save"); !ok {
		s.svcCtx.LocalCache.SetWithExpire("external_user_attr_save", 1, time.Hour*2)

		err = s.svcCtx.ModelExternalUserAttribute.DeleteByExternalUserId(s.ctx, externalUser.ExternalContact.ExternalUserID)

		if externalUser.ExternalContact.ExternalProfile != nil {
			for _, item := range externalUser.ExternalContact.ExternalProfile.ExternalAttr {

				ext := ""
				switch item.Type {
				case model.AttributeTypeText:
					ext, _ = jsoniter.MarshalToString(item.Text)
				case model.AttributeTypeWeb:
					ext, _ = jsoniter.MarshalToString(item.Web)
				case model.AttributeTypeMiniprogram:
					ext, _ = jsoniter.MarshalToString(item.MiniProgram)
				}

				externalUserAttrData := &model.TbExternalUserAttribute{
					ExternalUserid: externalUser.ExternalContact.ExternalUserID,
					AttributeType:  uint64(item.Type),
					AttributeValue: item.Name,
					Extension:      ext,
					Status:         1,
					CreatedAt:      ts,
					UpdatedAt:      ts,
				}
				_, _ = s.svcCtx.ModelExternalUserAttribute.Insert(s.ctx, externalUserAttrData)
			}
		}
	}

	// 外部用户信息关联信息保存
	// 没有关注信息结束 正常数据不会存在
	if externalUser.FollowInfo == nil {
		return nil
	}

	followData := &model.TbExternalUserFollow{
		ExternalUserid:    externalUser.ExternalContact.ExternalUserID,
		Unionid:           externalUser.ExternalContact.UnionID,
		Userid:            externalUser.FollowInfo.UserID,
		Platform:          crop,
		OperUserid:        externalUser.FollowInfo.OperUserID,
		AddWay:            uint64(externalUser.FollowInfo.AddWay),
		State:             externalUser.FollowInfo.State,
		StateChannel:      "",
		StateChannelValue: "",
		Remark:            externalUser.FollowInfo.Remark,
		RemarkMobiles:     strings.Join(externalUser.FollowInfo.RemarkMobiles, ","),
		Description:       externalUser.FollowInfo.Description,
		RemarkCorpName:    externalUser.FollowInfo.RemarkCorpName,
		RemarkPicMediaid:  "", // TODO 还没有此字段
		Status:            1,
		CreatedAt:         ts,
		UpdatedAt:         ts,
	}
	stateList := make([]string, 0)
	// "state": "99_9999#1#01234567890123456789"  或者 "99_9999#0#{"hid":15731}"
	// 充分校验下 防止正常渠道里带#号
	if strings.Contains(externalUser.FollowInfo.State, "#") {
		if stateList = strings.Split(followData.State, "#"); len(stateList) == 3 && utils.ArrayIn(stateList[1], []string{"0", "1"}) {
			followData.StateChannel = stateList[0]
			followData.StateChannelValue = stateList[2]
			// TODO 从缓存里面取
		}
	}

	err = retry.Do(func() error {
		dbExternalUserFollow, err = s.svcCtx.ModelExternalUserFollow.FindOneByExternalUserIdAndUserId(s.ctx, externalUser.ExternalContact.ExternalUserID, externalUser.FollowInfo.UserID, crop)
		return err
	}, retry.Attempts(3))
	if err != nil {
		return errors.New(fmt.Sprintf("获取旧外部用户关系信息失败, error: %v", err))
	}

	if dbExternalUserFollow != nil {
		followData.Seq = dbExternalUserFollow.Seq
		err = s.svcCtx.ModelExternalUserFollow.Update(s.ctx, dbExternalUserFollow)
	} else {
		_, err = s.svcCtx.ModelExternalUserFollow.Insert(s.ctx, followData)
	}

	if err != nil {
		return errors.New(fmt.Sprintf("保存外部用户关系信息失败, error: %v", err))
	}

	// 外部用户关系属性信息
	_ = s.svcCtx.ModelExternalUserFollowAttribute.DeleteByExternalUserIdAndUserIdAndPlatform(s.ctx, externalUser.ExternalContact.ExternalUserID, externalUser.FollowInfo.UserID, crop)
	for _, tag := range externalUser.FollowInfo.Tags {
		tagJson, _ := jsoniter.MarshalToString(tag)
		_, _ = s.svcCtx.ModelExternalUserFollowAttribute.Insert(s.ctx, &model.TbExternalUserFollowAttribute{
			ExternalUserid: externalUser.ExternalContact.ExternalUserID,
			Userid:         externalUser.FollowInfo.UserID,
			Platform:       crop,
			AttributeType:  model.AttributeTypeRemarkTag,
			AttributeValue: tag.TagID,
			Extension:      tagJson,
			Status:         1,
			CreatedAt:      ts,
			UpdatedAt:      ts,
		})
	}

	return nil
}
