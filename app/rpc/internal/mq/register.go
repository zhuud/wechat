package mq

import (
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/config"
	"rpc/internal/mq/externalcontactuser"
	"rpc/internal/svc"

	"github.com/spf13/cobra"
)

func RegisterMq(c config.Config, svcCtx *svc.ServiceContext) []*cobra.Command {
	defer func() {
		_ = logx.Close()
	}()

	return []*cobra.Command{
		{
			Use: "ConsumerSyncExternalUser",
			Run: externalcontactuser.NewSyncExternalUserConsumer(c.Kafka, svcCtx),
		},
	}
}
