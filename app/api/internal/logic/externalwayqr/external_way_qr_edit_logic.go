package externalwayqr

import (
	"context"
	"github.com/spf13/cast"
	"rpc/client/externalcontactway"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码更新
func NewExternalWayQrEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrEditLogic {
	return &ExternalWayQrEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrEditLogic) ExternalWayQrEdit(req *types.ExternalContactWayRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.ExternalcontactwayRpc.UpdateExternalContactWay(l.ctx, &externalcontactway.ExternalContactWayData{
		ConfigId:      req.ConfigID,
		Style:         req.Style,
		Remark:        req.Remark,
		SkipVerify:    req.SkipVerify,
		State:         req.State,
		User:          req.User,
		Party:         req.Party,
		ExpiresIn:     cast.ToInt32(req.ExpiresIn),
		ChatExpiresIn: cast.ToInt32(req.ChatExpiresIn),
		Unionid:       req.UnionID,
	})
	resp = &types.Response{
		Data: data,
	}

	return resp, err
}
