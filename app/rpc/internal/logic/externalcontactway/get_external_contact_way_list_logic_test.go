package externalcontactwaylogic

import (
	"context"
	"testing"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/davecgh/go-spew/spew"
)

func TestGetExternalContactWayListLogic_GetExternalContactWayList(t *testing.T) {

	r, e := NewGetExternalContactWayListLogic(context.Background(), svc.NewServiceContext()).GetExternalContactWayList(&wechat.ExternalContactWayListReq{
		Limit: 10,
	})
	spew.Dump(r, e)
}
