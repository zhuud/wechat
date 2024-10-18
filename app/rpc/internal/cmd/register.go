package cmd

import (
	"rpc/internal/cmd/externalcontactuser"
	"rpc/internal/cmd/externalcontactway"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterCmd(svcCtx *svc.ServiceContext) []*cobra.Command {
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
