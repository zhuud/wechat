package externalcontactway

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"rpc/internal/svc"
	"rpc/model"

	contactWayRequest "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	contactWayResponse "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/response"
	"github.com/avast/retry-go"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/fx"
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

	ctx := context.Background()

	getContactWayListFunc := func(source chan<- any) {

		_ = retry.Do(func() error {
			var err error
			params := &contactWayRequest.RequestListContactWay{
				Limit: 100,
			}
			list := &contactWayResponse.ResponseListContactWay{}
			list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(ctx, params)
			if err != nil {
				s.Error(err)
				return err
			}
			for len(list.ContactWayIDs) > 0 {
				for _, item := range list.ContactWayIDs {
					source <- item.ConfigID
				}

				params.Limit = 100
				params.Cursor = list.NextCursor
				list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(ctx, params)
				if err != nil {
					s.Error(err)
				}
			}
			return nil
		})
	}

	getContactWayInfoFunc := func(item any) any {
		configId, _ := item.(string)

		contactWayInfo, err := s.getContactWayInfo(ctx, configId)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("getContactWayInfo err: \n %v", err))
		}
		return contactWayInfo
	}

	saveContactWayInfo := func(item any) {
		contactWayInfo, _ := item.(*contactWayResponse.ResponseGetContactWay)
		err := s.saveContactWayInfo(ctx, contactWayInfo)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("saveContactWayInfo err: \n %v", err))
		}
	}

	// 主程序执行
	fx.From(getContactWayListFunc).Map(getContactWayInfoFunc).Parallel(saveContactWayInfo, fx.WithWorkers(64))

	return nil
}

//保存数据

func (s *syncExternalContactWayCmd) saveContactWayInfo(ctx context.Context, configData *contactWayResponse.ResponseGetContactWay) (err error) {
	err = retry.Do(func() error {
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
			ConfigId:      configData.ContactWay.ConfigID,
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

		findData, err := s.svcCtx.ModelUserServiceQrcodeModel.FindOneByConfigId(s.ctx, configData.ContactWay.ConfigID)
		if err != nil {
			s.Error(err)
		}
		if findData != nil {
			s.svcCtx.ModelUserServiceQrcodeModel.Update(s.ctx, userServiceQrcodeData)
		} else {
			s.svcCtx.ModelUserServiceQrcodeModel.Insert(s.ctx, userServiceQrcodeData)
		}
		return nil
	})
	return
}

// 获取详细信息
func (s *syncExternalContactWayCmd) getContactWayInfo(ctx context.Context, configId string) (configData *contactWayResponse.ResponseGetContactWay, err error) {
	err = retry.Do(func() error {
		configData, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.Get(context.Background(), configId)
		return nil
	})

	return
}

// 获取列表
func (s *syncExternalContactWayCmd) getContactWayList(ctx context.Context, params *contactWayRequest.RequestListContactWay, ch chan any) (err error) {

	err = retry.Do(func() error {
		list := &contactWayResponse.ResponseListContactWay{}
		list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(ctx, params)
		if err != nil {
			s.Error(err)
			return err
		}
		for len(list.ContactWayIDs) > 0 {
			for _, item := range list.ContactWayIDs {
				ch <- item.ConfigID
			}

			params.Limit = 100
			params.Cursor = list.NextCursor
			list, err = s.svcCtx.WeCom.WithCorp("yx").ContactWay.List(ctx, params)
			if err != nil {
				s.Error(err)
			}
		}
		return nil
	})

	return
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
