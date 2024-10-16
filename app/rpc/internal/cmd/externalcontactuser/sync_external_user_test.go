package externalcontactuser

import (
	"context"
	"testing"

	"rpc/internal/config"
	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
)

func Test(t *testing.T) {
	c := config.MustLoad()
	svcCtx := svc.NewServiceContext(c)
	err := newSyncExternalUserCmd(context.Background(), svcCtx).Do([]string{config.CropYx})
	spew.Dump(err)
}
