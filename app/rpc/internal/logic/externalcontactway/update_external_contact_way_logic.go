package externalcontactwaylogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExternalContactWayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExternalContactWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExternalContactWayLogic {
	return &UpdateExternalContactWayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateExternalContactWayLogic) UpdateExternalContactWay(in *wechat.ExternalContactWayData) (*wechat.SaveExternalContactWayResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.SaveExternalContactWayResp{}, nil
}
