package mq

import (
	"rpc/internal/config"
	"rpc/internal/mq/externalcontactuser"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterMq(c config.Config, svcCtx *svc.ServiceContext) []*cobra.Command {

	return []*cobra.Command{
		{
			Use: "ConsumerSyncExternalUser",
			Run: externalcontactuser.NewSyncExternalUserConsumer(c.Kafka, svcCtx),
		},
	}
}
