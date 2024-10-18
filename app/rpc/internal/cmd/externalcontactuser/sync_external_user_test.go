package externalcontactuser

import (
	"context"
	"testing"

	"rpc/internal/config"
	"rpc/internal/svc"

	"github.com/davecgh/go-spew/spew"
)

var ctx = context.Background()

func Test_SyncExternalUser(t *testing.T) {
	err := newSyncExternalUserCmd(ctx, svc.NewServiceContext()).Do([]string{config.CropYx})
	spew.Dump(err)
}

func Test_BizUser(t *testing.T) {
	r := svc.NewServiceContext().ModelBizUser.FindUidByUnionid(ctx, "oso0JtyDZ7qw7pGLseDN7m8ypiiM")
	ru := svc.NewServiceContext().ModelBizUser.FindUnionidByUid(ctx, 25913991167)
	spew.Dump(r, ru)
}
