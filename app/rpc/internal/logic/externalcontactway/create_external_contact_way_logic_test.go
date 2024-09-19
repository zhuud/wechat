package externalcontactwaylogic

import (
	"context"
	"fmt"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/davecgh/go-spew/spew"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"
	"testing"
)

func TestCreateExternalContactWayLogic_CreateExternalContactWay(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *wechat.ExternalContactWayData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *wechat.SaveExternalContactWayResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &CreateExternalContactWayLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.CreateExternalContactWay(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateExternalContactWay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateExternalContactWay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateExternalContactWay(t *testing.T) {
	var svcCtx *svc.ServiceContext = svc.NewServiceContext(config.MustLoad())

	req := &request2.RequestListContactWay{Limit: 10}
	list, err := svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), req)
	spew.Dump(err)
	fmt.Printf("%+#v", list)
}
