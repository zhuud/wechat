package main

import (
	"flag"

	"rpc/internal/cmd"
	"rpc/internal/mq"
	"rpc/internal/server"
	"rpc/internal/svc"

	"github.com/zhuud/go-library/svc/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	flag.Parse()

	svcCtx := svc.NewServiceContext()

	app.AddCommand(server.RegisterRpc(svcCtx))
	app.AddCommand(cmd.RegisterCmd(svcCtx)...)
	app.AddCommand(mq.RegisterMq(svcCtx)...)
	app.Run()
}
