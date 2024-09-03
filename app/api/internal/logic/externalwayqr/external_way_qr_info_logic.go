package externalwayqr

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
