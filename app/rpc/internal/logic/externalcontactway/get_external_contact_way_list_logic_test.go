package externalcontactwaylogic

import (
	"context"
	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var svcCtx *svc.ServiceContext

func init() {
	svcCtx = svc.NewServiceContext(config.MustLoad())
}

func TestGetExternalContactWayListLogic_GetExternalContactWayList(t *testing.T) {

	r, e := NewGetExternalContactWayListLogic(context.Background(), svcCtx).GetExternalContactWayList(&wechat.ExternalContactWayListReq{
		Limit: 10,
	})
	spew.Dump(r, e)
}
