package externalcontactwaylogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalContactWayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalContactWayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalContactWayListLogic {
	return &GetExternalContactWayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalContactWayListLogic) GetExternalContactWayList(in *wechat.ExternalContactWayListReq) (*wechat.ExternalContactWayListResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ExternalContactWayListResp{}, nil
}
