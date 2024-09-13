package externalcontactwaylogic

import (
	"context"
	"reflect"
	"rpc/internal/svc"
	"rpc/wechat"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestGetExternalContactWayInfoLogic_GetExternalContactWayInfo(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *wechat.ExternalContactWayReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *wechat.ExternalContactWayInfoResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GetExternalContactWayInfoLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.GetExternalContactWayInfo(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExternalContactWayInfoLogic.GetExternalContactWayInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExternalContactWayInfoLogic.GetExternalContactWayInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
