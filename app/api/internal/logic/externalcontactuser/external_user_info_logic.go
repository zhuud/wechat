package externalcontactuser

import (
    "context"

    "api/internal/svc"
    "api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type ExternalUserInfoLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 企微外部用户详情
func NewExternalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalUserInfoLogic {
    return &ExternalUserInfoLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ExternalUserInfoLogic) ExternalUserInfo(req *types.ExternalUserRequest) (resp *types.Response, err error) {

    // todo: add your logic here and delete this line
    

    return
}
