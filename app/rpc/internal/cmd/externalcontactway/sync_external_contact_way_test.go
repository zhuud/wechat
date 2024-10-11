package externalcontactway

import (
	"context"
	"testing"

	"rpc/internal/config"
	"rpc/internal/svc"
)

func Test(t *testing.T) {
	c := config.MustLoad()
	svcCtx := svc.NewServiceContext(c)
	newSyncExternalContactWayCmd(context.Background(), svcCtx).Do([]string{})
}