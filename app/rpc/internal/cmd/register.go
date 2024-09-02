package cmd

import (
	"github.com/zeromicro/go-zero/core/logx"
	"log"

	"rpc/internal/cmd/externalcontactuser"
	"rpc/internal/cmd/externalcontactway"
	"rpc/internal/config"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterCmd(c config.Config, svcCtx *svc.ServiceContext) []*cobra.Command {
	err := c.SetUp()
	if err != nil {
		log.Fatalf("cmd.RegisterCmd SetUp config:%v, error: %v", c, err)
	}
	defer func() {
		_ = logx.Close()
	}()

	return []*cobra.Command{
		{
			Use:  "CmdSyncExternalUser",
			RunE: externalcontactuser.NewSyncExternalUserCmd(svcCtx),
		},
		{
			Use:  "CmdSyncExternalContactWay",
			RunE: externalcontactway.NewSyncExternalContactWayCmd(svcCtx),
		},
	}
}
