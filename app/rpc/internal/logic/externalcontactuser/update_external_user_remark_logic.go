package externalcontactuserlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExternalUserRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExternalUserRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExternalUserRemarkLogic {
	return &UpdateExternalUserRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateExternalUserRemarkLogic) UpdateExternalUserRemark(in *wechat.UpdateExternalUserRemarkReq) (*wechat.ErrorResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ErrorResp{}, nil
}
