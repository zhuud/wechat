package externalcontactuserlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/logic/externalcontactuser/user"
	"rpc/internal/svc"
	"rpc/wechat"
)

type GetExternalUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserInfoLogic {
	return &GetExternalUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalUserInfoLogic) GetExternalUserInfo(in *wechat.ExternalUserInfoReq) (resp *wechat.ExternalUserInfoResp, err error) {
	// 基础信息
	userList, err := user.NewGetExternalUserCacheLogic(l.ctx, l.svcCtx).GetUserCache(in)
	if err != nil || userList == nil || len(userList) == 0 {
		return
	}

	// uint信息
	userUnit := user.NewGetExternalUserUnitLogic(l.ctx, l.svcCtx).GetUserUint(userList, in)

	resp.List = map[string]*wechat.ExternalUserInfo{}
	for _, externalUserId := range in.ExternalUseridList {
		if _, ok := userList[externalUserId]; !ok {
			continue
		}

		userInfo := userList[externalUserId]

		//关注的用户信息
		user.NewGetExternalUserFormatLogic(l.ctx, l.svcCtx).HandleFollowUserFormat(userInfo, userUnit)

		resp.List[externalUserId] = &wechat.ExternalUserInfo{
			ExternalUser: userInfo.ExternalUser,
			FollowUser:   userInfo.ExternalUserFollow,
		}
	}

	return resp, nil
}
