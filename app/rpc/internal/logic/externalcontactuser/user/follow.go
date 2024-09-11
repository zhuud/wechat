package user

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/wechat"
)

type GetExternalUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserFollowLogic {
	return &GetExternalUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 外部联系人属性表
 */
func (t *GetExternalUserFollowLogic) GetUserFollowListByExternalUserIdList(externalUserIdList []string) (externalUserListMap map[string][]*wechat.ExternalUserFollowUser, err error) {
	externalUserFollowList, err := t.svcCtx.ModelExternalUserFollow.FindListByExternalUserId(t.ctx, externalUserIdList)
	if err != nil {
		return externalUserListMap, err
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

	return
}
