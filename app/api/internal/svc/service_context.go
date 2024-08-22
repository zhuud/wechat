package svc

import (
    "api/internal/config"
    "rpc/client/externalcontactuser"

    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config                 config.Config
    ExternalContactUserRpc externalcontactuser.ExternalContactUser
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config: c,

        ExternalContactUserRpc: externalcontactuser.NewExternalContactUser(zrpc.MustNewClient(c.WechatRpc)),
    }
}
