package externalwayqr

import (
	"context"
	"rpc/client/externalcontactway"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码删除
func NewExternalWayQrDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrDeleteLogic {
	return &ExternalWayQrDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrDeleteLogic) ExternalWayQrDelete(req *types.ExternalContactWayDelRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.ExternalcontactwayRpc.DeleteExternalContactWay(l.ctx, &externalcontactway.ExternalContactWayReq{
		ConfigId: req.ConfigID,
	})
	resp = &types.Response{
		Data: data,
	}

	return resp, err
}
