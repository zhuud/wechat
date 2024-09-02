package externalcontactway

import (
	"context"

	"rpc/internal/svc"
	"rpc/model"

	contactWayRequest "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	contactWayResponse "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/response"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewSyncExternalContactWayCmd(svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return newSyncExternalContactWayCmd(cmd.Context(), svcCtx).Do(args)
	}
}

type syncExternalContactWayCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func newSyncExternalContactWayCmd(ctx context.Context, svcCtx *svc.ServiceContext) *syncExternalContactWayCmd {
	return &syncExternalContactWayCmd{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *syncExternalContactWayCmd) Do(args []string) error {

	params := &contactWayRequest.RequestListContactWay{}
	params.Limit = 100

	var err error
	list := &contactWayResponse.ResponseListContactWay{}

	list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)
	if err != nil {
		s.Error(err)
		return err
	}
	for len(list.ContactWayIDs) > 0 {
		for _, item := range list.ContactWayIDs {

			configId := item.ConfigID
			configData, err := s.svcCtx.WeCom.WithCorp("yx").ContactWay.Get(context.Background(), configId)
			if err != nil || configData.ContactWay == nil {
				s.Error(err)
				continue
			}
			s.svcCtx.ModelUserServiceQrcodeModel.Insert(s.ctx, &model.UserServiceQrcode{
				ConfigId:  configId,
				Type:      int64(configData.ContactWay.Type),
				Scene:     int64(configData.ContactWay.Scene),
				State:     configData.ContactWay.State,
				QrCode:    configData.ContactWay.QrCode,
				Remark:    configData.ContactWay.Remark,
				Style:     int64(configData.ContactWay.Style),
				ExpiresIn: int64(configData.ContactWay.ExpiresIn),
			})
		}

		params.Limit = 100
		params.Cursor = list.NextCursor
		list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)
		if err != nil {
			s.Error(err)
		}
	}

	return nil
}
