package externalcontactuser

import (
	"context"
	"testing"

	"rpc/internal/config"
	"rpc/internal/svc"
)

func Test(t *testing.T) {
	c := config.MustLoad()
	svcCtx := svc.NewServiceContext(c)
	newSyncExternalUserCmd(context.Background(), svcCtx).Do([]string{"yx"})
}
