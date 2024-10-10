package externalcontactuser

import (
	"context"
	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuud/go-library/svc/kafka"
)

const syncExternalUserTopic = "5002,5003"

func NewSyncExternalUserConsumer(kafkaConf kq.KqConf, svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		kafkaConf.Topic = syncExternalUserTopic
		kafka.Consume(kafkaConf, cmd.Use, newSyncExternalUserConsumer(cmd.Context(), svcCtx))
	}
}

type syncExternalUserConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func newSyncExternalUserConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *syncExternalUserConsumer {
	return &syncExternalUserConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *syncExternalUserConsumer) Consume(ctx context.Context, key string, value string) error {

	// do code ...
	spew.Dump(key, value)

	return nil
}
