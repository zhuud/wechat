package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
)

type GetExternalUserTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserTagLogic {
	return &GetExternalUserTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
