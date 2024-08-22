package main

import (
	"flag"
	"fmt"

	"rpc/internal/config"
	externalcontactgroupchatserver "rpc/internal/server/externalcontactgroupchat"
	externalcontactuserserver "rpc/internal/server/externalcontactuser"
	externalcontactwayserver "rpc/internal/server/externalcontactway"
	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	c := config.MustLoad(*configFile)
	svcCtx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		wechat.RegisterExternalContactUserServer(grpcServer, externalcontactuserserver.NewExternalContactUserServer(svcCtx))
		wechat.RegisterExternalContactGroupChatServer(grpcServer, externalcontactgroupchatserver.NewExternalContactGroupChatServer(svcCtx))
		wechat.RegisterExternalContactWayServer(grpcServer, externalcontactwayserver.NewExternalContactWayServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting Rpc Server At %s...\n", c.ListenOn)
	s.Start()
}
