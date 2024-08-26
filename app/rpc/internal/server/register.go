package server

import (
	"fmt"

	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	externalcontactgroupchatserver "rpc/internal/server/externalcontactgroupchat"
	externalcontactuserserver "rpc/internal/server/externalcontactuser"
	externalcontactwayserver "rpc/internal/server/externalcontactway"
)

func RegisterRpc(c config.Config, svcCtx *svc.ServiceContext) *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {

			s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
				register(grpcServer, svcCtx)

				if c.Mode == service.DevMode || c.Mode == service.TestMode {
					reflection.Register(grpcServer)
				}
			})
			defer s.Stop()

			fmt.Printf("Starting Rpc Server At %s...\n", c.ListenOn)
			s.Start()
		},
	}
}

func register(grpcServer *grpc.Server, svcCtx *svc.ServiceContext) {
	wechat.RegisterExternalContactUserServer(grpcServer, externalcontactuserserver.NewExternalContactUserServer(svcCtx))
	wechat.RegisterExternalContactGroupChatServer(grpcServer, externalcontactgroupchatserver.NewExternalContactGroupChatServer(svcCtx))
	wechat.RegisterExternalContactWayServer(grpcServer, externalcontactwayserver.NewExternalContactWayServer(svcCtx))
}
