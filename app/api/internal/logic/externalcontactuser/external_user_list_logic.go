package externalcontactuser

import (
    "api/internal/svc"
    "api/internal/types"
    "context"
    "rpc/client/externalcontactuser"

    "github.com/zeromicro/go-zero/core/logx"
)

type ExternalUserListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 企微外部用户详情列表
func NewExternalUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalUserListLogic {
    return &ExternalUserListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ExternalUserListLogic) ExternalUserList(req *types.ExternalUserRequest) (resp *types.Response, err error) {
    data, err := l.svcCtx.ExternalContactUserRpc.GetExternalUserInfo(l.ctx, &externalcontactuser.ExternalUserInfoReq{
        ExternalUseridList: req.ExternalUseridList,
        UnionidList:        req.UnionidList,
    })
    resp = &types.Response{
        Data: data,
    }
    return resp, err
}
