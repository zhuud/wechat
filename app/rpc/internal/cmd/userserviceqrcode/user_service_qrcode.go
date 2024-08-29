package userserviceqrcode

import (
	"context"
	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserServiceQrcodeCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserServiceQrcodeCmd(svcCtx *svc.ServiceContext) *UserServiceQrcodeCmd {
	return &UserServiceQrcodeCmd{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceQrcodeCmd) WithCtx(ctx context.Context) {
	s.ctx = ctx
	s.Logger = logx.WithContext(s.ctx)
}

func (s *UserServiceQrcodeCmd) UserServiceQrcode(c *cobra.Command, args []string) error {
	s.WithCtx(c.Context())

	// do code ...
	spew.Dump(args)

	return nil
}