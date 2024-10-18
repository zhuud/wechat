package mq

import (
	"rpc/internal/mq/externalcontactuser"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterMq(svcCtx *svc.ServiceContext) []*cobra.Command {
	return []*cobra.Command{
		{
			Use: "ConsumerSyncExternalUser",
			Run: externalcontactuser.NewSyncExternalUserConsumer(svcCtx.Config.Kafka, svcCtx),
		},
	}
}
