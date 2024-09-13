package externalwayqr

import (
	"context"
	"rpc/client/externalcontactway"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码列表
func NewExternalWayQrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrListLogic {
	return &ExternalWayQrListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrListLogic) ExternalWayQrList(req *types.ExternalWayQrListRequest) (resp *types.Response, err error) {

	data, err := l.svcCtx.ExternalcontactwayRpc.GetExternalContactWayList(l.ctx, &externalcontactway.ExternalContactWayListReq{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Limit:     10,
		Cursor:    req.Cursor,
	})
	resp = &types.Response{
		Data: data,
	}
	return resp, err

}
