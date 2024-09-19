package svc

import (
    "api/internal/config"
    "rpc/client/externalcontactuser"

    "rpc/client/externalcontactway"
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config                 config.Config
    ExternalContactUserRpc externalcontactuser.ExternalContactUser
    ExternalcontactwayRpc  externalcontactway.ExternalContactWay
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config: c,
        ExternalContactUserRpc: externalcontactuser.NewExternalContactUser(zrpc.MustNewClient(c.WechatRpc)),
        ExternalcontactwayRpc:  externalcontactway.NewExternalContactWay(zrpc.MustNewClient(c.WechatRpc)),
    }
}
