package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/wechat"
)

type GetExternalUserBaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserBaseLogic {
	return &GetExternalUserBaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 获取用户基础信息
 */
func (t *GetExternalUserBaseLogic) GetUserListByExternalUserIdList(externalUserIdList []string) (externalUserMap map[string]*wechat.ExternalUser, err error) {
	externalUserList, err := t.svcCtx.ModelExternalUser.FindListByExternalUserid(t.ctx, externalUserIdList)
	if err != nil {
		return nil, err
	}

	externalUser := make(map[string]*wechat.ExternalUser)
	for _, user := range externalUserList {
		externalUser[user.ExternalUserid] = &wechat.ExternalUser{
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
	}

	return
}
