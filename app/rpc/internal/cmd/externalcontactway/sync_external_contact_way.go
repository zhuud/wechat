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
			/**
					Id            uint64    `db:"id"`              // 主键id
			ConfigId      string    `db:"config_id"`       // 新增联系方式的配置id
			Type          int64     `db:"type"`            // 联系方式类型，1-单人，2-多人
			Scene         int64     `db:"scene"`           // 场景，1-在小程序中联系，2-通过二维码联系
			Style         int64     `db:"style"`           // 小程序中联系按钮的样式，仅在scene为1时返回，详见附录
			Remark        string    `db:"remark"`          // 联系方式的备注信息，用于助记
			SkipVerify    int64     `db:"skip_verify"`     // 外部客户添加时是否无需验证 0-否 1-是
			State         string    `db:"state"`           // 企业自定义的state参数，用于区分不同的添加渠道
			QrCode        string    `db:"qr_code"`         // 联系二维码的URL，仅在scene为2时返回
			User          string    `db:"user"`            // 使用该联系方式的用户userID列表
			Party         string    `db:"party"`           // 使用该联系方式的部门id列表
			IsTemp        int64     `db:"is_temp"`         // 是否临时会话模式0 不是 1 是
			ExpiresIn     int64     `db:"expires_in"`      // 临时会话二维码有效期，以秒为单位
			ChatExpiresIn int64     `db:"chat_expires_in"` // 临时会话有效期，以秒为单位
			Unionid       string    `db:"unionid"`         // 可进行临时会话的客户unionid
			IsExclusive   int64     `db:"is_exclusive"`    // 0-否 1-是；是否开启同一外部企业客户只能添加同一个员工**/

			var skipVerifyIntValeue int64
			if configData.ContactWay.SkipVerify {
				skipVerifyIntValeue = 1
			}

			var isTemp int64
			if configData.ContactWay.IsTemp {
				isTemp = 1
			}

			var isExclusive int64
			
			userList,_ := json.Marshal(configData.ContactWay.User)
			partyList,_ := json.Marshal(configData.ContactWay.Party)

			s.svcCtx.ModelUserServiceQrcodeModel.Insert(s.ctx, &model.UserServiceQrcode{
				ConfigId:  configId,
				Type:      int64(configData.ContactWay.Type),
				Scene:     int64(configData.ContactWay.Scene),
				State:     configData.ContactWay.State,
				QrCode:    configData.ContactWay.QrCode,
				Remark:    configData.ContactWay.Remark,
				Style:     int64(configData.ContactWay.Style),
				ExpiresIn: int64(configData.ContactWay.ExpiresIn),
				SkipVerify: skipVerifyIntValeue,
				User: string(userList),
				Party: string(partyList),
				IsTemp: isTemp,
				ChatExpiresIn: int64(configData.ContactWay.ChatExpiresIn),
				Unionid: configData.ContactWay.UnionID,
				IsExclusive: isExclusive,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
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
