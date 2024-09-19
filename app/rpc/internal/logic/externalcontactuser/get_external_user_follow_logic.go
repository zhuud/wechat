package externalcontactuserlogic

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/logic/externalcontactuser/user"
	"rpc/internal/svc"
	"rpc/wechat"
)

type GetExternalUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserInfoLogic {
	return &GetExternalUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (t *GetExternalUserFollowLogic) GetUserFollowByFollowIdList(userIdList []string) (list map[string]*wechat.ExternalUserFollowUser) {
	if len(userIdList) == 0 {
		return
	}

	followList, err := t.svcCtx.ModelExternalUserFollow.FindListByExternalUserId(t.ctx, userIdList)
	if err != nil {
		t.Logger.Error(`FindListByExternalUserId_err`, err)
		return
	}
	for _, follow := range followList {
		mobileList := make([]string, 0)
		_ = jsoniter.UnmarshalFromString(follow.RemarkMobiles, &mobileList)
		list[follow.Userid] = &wechat.ExternalUserFollowUser{
			Userid:         follow.Userid,
			Remark:         follow.Remark,
			Description:    follow.Description,
			Createtime:     int32(follow.CreatedAt.Unix()),
			Tags:           nil, // TODO
			RemarkCorpName: follow.RemarkCorpName,
			RemarkMobiles:  mobileList,
			OperUserid:     follow.OperUserid,
			AddWay:         int32(follow.AddWay),
			WechatChannels: nil, // TODO

		}
	}

	return
}

func (t *GetExternalUserFollowLogic) GetUserFollowByUserIdList(externalUserIdList []string) (list map[string][]*wechat.ExternalUserFollowUser) {
	if len(externalUserIdList) == 0 {
		return
	}

	list, err := user.NewGetExternalUserFollowLogic(t.ctx, t.svcCtx).GetUserFollowListByExternalUserIdList(externalUserIdList)
	if err != nil {
		t.Logger.Error(err, `GetUserFollowListByExternalUserIdList_err`)
	}
	return
}
