package externalcontactuserlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalUserIdByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserIdByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserIdByUserIdLogic {
	return &GetExternalUserIdByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalUserIdByUserIdLogic) GetExternalUserIdByUserId(in *wechat.ExternalUserIdReq) (*wechat.ExternalUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ExternalUserIdResp{}, nil
}
