package externalcontactway

import (
	"context"
	"encoding/json"
	"time"

	"rpc/internal/svc"
	"rpc/model"

	contactWayRequest "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	contactWayResponse "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/response"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
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

	limiter := rate.NewLimiter(10, 10)

	return s.ContactWayTask(params, limiter)
}

func (s *syncExternalContactWayCmd) ContactWayTask(params *contactWayRequest.RequestListContactWay, limiter *rate.Limiter) (err error) {

	list := &contactWayResponse.ResponseListContactWay{}

	limiter.Wait(s.ctx)
	list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)
	if err != nil {
		s.Error(err)
		return err
	}
	for len(list.ContactWayIDs) > 0 {
		for _, item := range list.ContactWayIDs {

			configId := item.ConfigID

			limiter.Wait(s.ctx)
			configData, err := s.svcCtx.WeCom.WithCorp("yx").ContactWay.Get(context.Background(), configId)
			if err != nil || configData.ContactWay == nil {
				s.Error(err)
				continue
			}
			var skipVerifyIntValeue int64
			if configData.ContactWay.SkipVerify {
				skipVerifyIntValeue = 1
			}

			var isTemp int64
			if configData.ContactWay.IsTemp {
				isTemp = 1
			}

			var isExclusive int64

			userList, _ := json.Marshal(configData.ContactWay.User)
			partyList, _ := json.Marshal(configData.ContactWay.Party)

			userServiceQrcodeData := &model.UserServiceQrcode{
				ConfigId:      configId,
				Type:          int64(configData.ContactWay.Type),
				Scene:         int64(configData.ContactWay.Scene),
				State:         configData.ContactWay.State,
				QrCode:        configData.ContactWay.QrCode,
				Remark:        configData.ContactWay.Remark,
				Style:         int64(configData.ContactWay.Style),
				ExpiresIn:     int64(configData.ContactWay.ExpiresIn),
				SkipVerify:    skipVerifyIntValeue,
				User:          string(userList),
				Party:         string(partyList),
				IsTemp:        isTemp,
				ChatExpiresIn: int64(configData.ContactWay.ChatExpiresIn),
				Unionid:       configData.ContactWay.UnionID,
				IsExclusive:   isExclusive,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}

			limiter.Wait(s.ctx)
			findData, err := s.svcCtx.ModelUserServiceQrcodeModel.FindOneByConfigId(s.ctx, configId)
			if err != nil {
				s.Error(err)
			}
			if findData != nil {
				s.svcCtx.ModelUserServiceQrcodeModel.Update(s.ctx, userServiceQrcodeData)
			} else {
				s.svcCtx.ModelUserServiceQrcodeModel.Insert(s.ctx, userServiceQrcodeData)
			}
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
