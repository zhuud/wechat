package externalwayqr

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
