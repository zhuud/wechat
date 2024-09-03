package externalcontactuser

import (
	"context"
	"time"

	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewSyncExternalUserCmd(svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return newSyncExternalUserCmd(cmd.Context(), svcCtx).Do(args)
	}
}

type syncExternalUserCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func newSyncExternalUserCmd(ctx context.Context, svcCtx *svc.ServiceContext) *syncExternalUserCmd {
	return &syncExternalUserCmd{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *syncExternalUserCmd) Do(args []string) error {
	st := time.Now()
	logx.WithDuration(time.Since(st)).Info("hhh")
	// do code ...
	spew.Dump(args)
	return nil
}
