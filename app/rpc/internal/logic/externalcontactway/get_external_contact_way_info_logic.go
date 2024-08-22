package externalcontactwaylogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalContactWayInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalContactWayInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalContactWayInfoLogic {
	return &GetExternalContactWayInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalContactWayInfoLogic) GetExternalContactWayInfo(in *wechat.ExternalContactWayReq) (*wechat.ExternalContactWayInfoResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ExternalContactWayInfoResp{}, nil
}
