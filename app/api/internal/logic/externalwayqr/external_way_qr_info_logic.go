package externalwayqr

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"rpc/client/externalcontactway"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码详情
func NewExternalWayQrInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrInfoLogic {
	return &ExternalWayQrInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrInfoLogic) ExternalWayQrInfo(req *types.ExternalWayQrInfoRequest) (resp *types.Response, err error) {
	data, err := l.svcCtx.ExternalcontactwayRpc.GetExternalContactWayInfo(l.ctx, &externalcontactway.ExternalContactWayReq{
		ConfigId: req.ConfigId,
	})
	resp = &types.Response{
		Data: data,
	}
	return resp, err
}
