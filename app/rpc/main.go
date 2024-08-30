package main

import (
	"rpc/internal/cmd"
	"rpc/internal/config"
	"rpc/internal/mq"
	"rpc/internal/server"
	"rpc/internal/svc"

	"github.com/zhuud/go-library/svc/app"
	_ "go.uber.org/automaxprocs"
)

func main() {

	c := config.MustLoad()
	svcCtx := svc.NewServiceContext(c)

	app.AddCommand(server.RegisterRpc(c, svcCtx))
	app.AddCommand(cmd.RegisterCmd(c, svcCtx)...)
	app.AddCommand(mq.RegisterMq(c, svcCtx)...)
	app.Run()
}
