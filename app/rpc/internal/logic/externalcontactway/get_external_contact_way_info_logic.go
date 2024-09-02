package externalcontactwaylogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

var svcCtx *svc.ServiceContext

func init() {
	svcCtx = svc.NewServiceContext(config.MustLoad())
}

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

	configData, err := svcCtx.WeCom.WithCorp("yx").ContactWay.Get(context.Background(), in.ConfigId)
	
	if err != nil {
		return nil, err
	}

	return &wechat.ExternalContactWayInfoResp{
		ContactWay: &wechat.ExternalContactWayData{
			ConfigId: configData.ContactWay.ConfigID,
			QrCode: configData.ContactWay.QrCode,
			
		},
	}, nil
}
