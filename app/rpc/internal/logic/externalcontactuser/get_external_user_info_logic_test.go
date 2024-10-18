package externalcontactuserlogic

import (
	"context"
	"reflect"
	"testing"

	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"

	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/davecgh/go-spew/spew"
)

func Test_GetExternalUserInfo(t *testing.T) {
	type args struct {
		param *wechat.ExternalUserInfoReq
	}
	tests := []struct {
		name    string
		args    args
		want    *wechat.ExternalUserInfoResp
		wantErr error
	}{
		{
			name: "",
			args: args{
				param: &wechat.ExternalUserInfoReq{
					UnionidList:        []string{"wmYYltDAAAlg093GN65jtwLAn1VqOi5g"},
					ExternalUseridList: []string{`wmYYltDAAAlg093GN65jtwLAn1VqOi5g`},
				},
			},
			want:    &wechat.ExternalUserInfoResp{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewGetExternalUserInfoLogic(context.Background(), svc.NewServiceContext())
			if got, err := l.GetExternalUserInfo(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("res = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}

	r, e := NewGetExternalUserInfoLogic(context.Background(), svc.NewServiceContext()).GetExternalUserInfo(&wechat.ExternalUserInfoReq{
		ExternalUseridList: []string{"wmYYltDAAAlg093GN65jtwLAn1VqOi5g", "sssss"},
	})
	spew.Dump(r, e)
	return
}

func Test_Get(t *testing.T) {
	spew.Dump(svc.NewServiceContext().WeCom.WithCorp(config.CropYx).ExternalUser.Get(context.Background(), "wmYYltDAAAlg093GN65jtwLAn1VqOi5g", ""))
	return
}

func Test_GetByUserId(t *testing.T) {
	prasms := &request2.RequestListContactWay{Limit: 10}
	list, err := svc.NewServiceContext().WeCom.WithCorp(config.CropYx).ContactWay.List(context.Background(), prasms)
	spew.Dump(list, err)
}
