package externalcontactwaylogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"rpc/internal/svc"
	"rpc/wechat"
	"testing"
)

func TestUpdateExternalContactWayLogic_UpdateExternalContactWay(t *testing.T) {
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
			l := &UpdateExternalContactWayLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.UpdateExternalContactWay(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateExternalContactWay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateExternalContactWay() got = %v, want %v", got, tt.want)
			}
		})
	}
}
