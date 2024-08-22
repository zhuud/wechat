package externalcontactuserlogic

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"rpc/internal/svc"
	"rpc/wechat"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type (
	externalUserUnit struct {
		User           sync.Map
		UserFollow     sync.Map
		UserFollowAttr sync.Map
	}
)

func NewGetExternalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserInfoLogic {
	return &GetExternalUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalUserInfoLogic) GetExternalUserInfo(in *wechat.ExternalUserInfoReq) (*wechat.ExternalUserInfoResp, error) {

	externalUserList, err := l.svcCtx.ModelExternalUser.FindListByExternalUserid(l.ctx, in.ExternalUseridList)
	if err != nil {
		return nil, err
	}
	externalUserFollowList, err := l.svcCtx.ModelExternalUserFollow.FindListByExternalUserid(l.ctx, in.ExternalUseridList)
	if err != nil {
		return nil, err
	}

	externalUserFollowCard := make(map[string][]*wechat.ExternalUserFollowUser)
	for _, follow := range externalUserFollowList {
		mobileList := make([]string, 0)
		_ = jsoniter.UnmarshalFromString(follow.RemarkMobiles, &mobileList)
		externalUserFollowCard[follow.ExternalUserid] = append(externalUserFollowCard[follow.ExternalUserid], &wechat.ExternalUserFollowUser{
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
		})
	}

	externalUserCard := make(map[string]*wechat.ExternalUserInfo)
	for _, user := range externalUserList {
		externalUserCard[user.ExternalUserid] = &wechat.ExternalUserInfo{}
		externalUserCard[user.ExternalUserid].ExternalUser = &wechat.ExternalUser{
			ExternalUserid:  user.ExternalUserid,
			Name:            user.Name,
			Position:        user.Position,
			Avatar:          user.Avatar,
			CorpName:        user.CorpName,
			CorpFullName:    user.CorpFullName,
			Type:            int32(user.Type),
			Gender:          int32(user.Gender),
			Unionid:         user.Unionid,
			ExternalProfile: nil,
		}
		if follow, ok := externalUserFollowCard[user.ExternalUserid]; ok {
			externalUserCard[user.ExternalUserid].FollowUser = follow
		}
	}

	return &wechat.ExternalUserInfoResp{
		List: externalUserCard,
	}, nil
}
