package externalwayqr

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码添加
func NewExternalWayQrAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrAddLogic {
	return &ExternalWayQrAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrAddLogic) ExternalWayQrAdd(req *types.ExternalContactWayRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
