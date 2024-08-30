package cmd

import (
	"rpc/internal/cmd/externalcontactuser"
	"rpc/internal/cmd/userserviceqrcode"
	"rpc/internal/config"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterCmd(_ config.Config, svcCtx *svc.ServiceContext) []*cobra.Command {
	return []*cobra.Command{
		{
			Use:  "CmdSyncExternalUser",
			RunE: externalcontactuser.NewSyncExternalUserCmd(svcCtx).SyncExternalUser,
		},
		{
			Use:  "CmdUserServiceQrcodeSync",
			RunE: userserviceqrcode.NewUserServiceQrcodeCmd(svcCtx).UserServiceQrcode,
		},
	}
}
