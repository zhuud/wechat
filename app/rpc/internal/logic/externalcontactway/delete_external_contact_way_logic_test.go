package externalcontactwaylogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"testing"
)

func TestDeleteExternalContactWayLogic_DeleteExternalContactWay(t *testing.T) {
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
		want    *wechat.ErrorResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DeleteExternalContactWayLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.DeleteExternalContactWay(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteExternalContactWay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteExternalContactWay() got = %v, want %v", got, tt.want)
			}
		})
	}
}
