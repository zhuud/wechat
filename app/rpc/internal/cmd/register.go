package cmd

import (
	"log"

	"rpc/internal/cmd/externalcontactuser"
	"rpc/internal/config"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterCmd(c config.Config, svcCtx *svc.ServiceContext) []*cobra.Command {
	err := c.SetUp()
	if err != nil {
		log.Fatalf("cmd.RegisterCmd SetUp config:%v, error: %v", c, err)
	}

	return []*cobra.Command{
		{
			Use:  "CmdSyncExternalUser",
			RunE: externalcontactuser.NewSyncExternalUserCmd(svcCtx),
		},
	}
}
