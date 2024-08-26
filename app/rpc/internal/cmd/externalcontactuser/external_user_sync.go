package externalcontactuser

import (
	"context"
	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncExternalUserCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncExternalUserCmd(svcCtx *svc.ServiceContext) *SyncExternalUserCmd {
	return &SyncExternalUserCmd{
		svcCtx: svcCtx,
	}
}

func (s *SyncExternalUserCmd) WithCtx(ctx context.Context) {
	s.ctx = ctx
	s.Logger = logx.WithContext(s.ctx)
}

func (s *SyncExternalUserCmd) SyncExternalUser(c *cobra.Command, args []string) error {
	s.WithCtx(c.Context())

	// do code ...
	spew.Dump(args)

	return nil
}
