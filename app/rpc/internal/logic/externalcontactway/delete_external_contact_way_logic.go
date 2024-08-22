package externalcontactwaylogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExternalContactWayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteExternalContactWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExternalContactWayLogic {
	return &DeleteExternalContactWayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteExternalContactWayLogic) DeleteExternalContactWay(in *wechat.ExternalContactWayReq) (*wechat.ErrorResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ErrorResp{}, nil
}
