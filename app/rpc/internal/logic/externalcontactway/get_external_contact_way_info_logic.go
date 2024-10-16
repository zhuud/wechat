package externalcontactwaylogic

import (
	"context"

	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalContactWayInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalContactWayInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalContactWayInfoLogic {
	return &GetExternalContactWayInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalContactWayInfoLogic) GetExternalContactWayInfo(in *wechat.ExternalContactWayReq) (*wechat.ExternalContactWayInfoResp, error) {
	// todo: add your logic here and delete this line

	configData, err := l.svcCtx.WeCom.WithCorp(config.CropYx).ContactWay.Get(context.Background(), in.ConfigId)

	if err != nil {
		return nil, err
	}

	return &wechat.ExternalContactWayInfoResp{
		ContactWay: &wechat.ExternalContactWayData{
			ConfigId:      configData.ContactWay.ConfigID,
			QrCode:        configData.ContactWay.QrCode,
			Type:          int32(configData.ContactWay.Type),
			Scene:         int32(configData.ContactWay.Scene),
			User:          configData.ContactWay.User,
			IsTemp:        configData.ContactWay.IsTemp,
			ExpiresIn:     int32(configData.ContactWay.ExpiresIn),
			ChatExpiresIn: int32(configData.ContactWay.ChatExpiresIn),
			Unionid:       configData.ContactWay.UnionID,
			Conclusions:   &wechat.ExternalContactWayConclusion{},
		},
	}, nil
}
