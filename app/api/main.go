package main

import (
	"flag"
	"fmt"
	"net/http"

	"api/internal/config"
	"api/internal/handler"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	_ "go.uber.org/automaxprocs"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	c := config.MustLoad(*configFile)
	svcCtx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, svcCtx)

	httpx.SetErrorHandler(func(err error) (int, any) {
		return http.StatusOK, types.Response{
			Code: 1000,
			Msg:  err.Error(),
		}
	})

	fmt.Printf("Starting Http Server At %s:%d...\n", c.Host, c.Port)
	server.Start()
}
