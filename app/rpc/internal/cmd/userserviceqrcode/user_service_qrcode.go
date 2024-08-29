package userserviceqrcode

import (
	"context"
	"rpc/internal/svc"
	"rpc/model"

	"rpc/internal/config"

	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/response"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

var svcCtx *svc.ServiceContext

func init() {
	svcCtx = svc.NewServiceContext(config.MustLoad())
}

type UserServiceQrcodeCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserServiceQrcodeCmd(svcCtx *svc.ServiceContext) *UserServiceQrcodeCmd {
	return &UserServiceQrcodeCmd{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceQrcodeCmd) WithCtx(ctx context.Context) {
	s.ctx = ctx
	s.Logger = logx.WithContext(s.ctx)
}

func (s *UserServiceQrcodeCmd) UserServiceQrcode(c *cobra.Command, args []string) error {
	s.WithCtx(c.Context())

	params := &request2.RequestListContactWay{}
	params.Limit = 100

	var err error
	list := &response3.ResponseListContactWay{}

	list, err = svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)
	if err != nil {
		logx.Error(err)
		return err
	}
	for len(list.ContactWayIDs) > 0 {
		for _, item := range list.ContactWayIDs {

			configId := item.ConfigID
			configData,err := svcCtx.WeCom.WithCorp("yx").ContactWay.Get(context.Background(), configId)
			if err != nil || configData.ContactWay == nil{
				logx.Error(err)
				continue
			}
			svcCtx.ModelUserServiceQrcodeModel.Insert(s.ctx, &model.UserServiceQrcode{
				ConfigId:  configId,
				Type: int64(configData.ContactWay.Type),
				Scene: int64(configData.ContactWay.Scene),
				State: configData.ContactWay.State,
				QrCode: configData.ContactWay.QrCode,
				Remark: configData.ContactWay.Remark,
				Style: int64(configData.ContactWay.Style),
				ExpiresIn: int64(configData.ContactWay.ExpiresIn),
			})
		}

		params.Limit = 100
		params.Cursor = list.NextCursor
		list, err = svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)
		if err != nil {
			logx.Error(err)
		}
	}
	
	return nil
}
