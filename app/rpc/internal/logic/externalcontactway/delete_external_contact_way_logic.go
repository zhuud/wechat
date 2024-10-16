package externalcontactwaylogic

import (
	"context"
	"errors"

	"rpc/internal/config"
	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExternalContactWayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteExternalContactWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExternalContactWayLogic {
	return &DeleteExternalContactWayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteExternalContactWayLogic) DeleteExternalContactWay(in *wechat.ExternalContactWayReq) (*wechat.ErrorResp, error) {
	// todo: add your logic here and delete this line

	if in.ConfigId == "" {
		return nil, errors.New("参数错误-config_id")
	}

	// 调用企微删除企业已配置的「联系我」方式接口
	_, err := l.svcCtx.WeCom.WithCorp(config.CropYx).ContactWay.Delete(l.ctx, in.ConfigId)
	if err != nil {
		l.Logger.Error("ContactWay_Delete_Err", err)
		return nil, err
	}

	// 本地结构化-更新
	qrcodeInfo, findErr := l.svcCtx.ModelUserServiceQrcodeModel.FindOneByConfigId(l.ctx, in.ConfigId)
	if findErr != nil {
		l.Logger.Error("ModelUserServiceQrcodeModel_FindOneByConfigId_Err", findErr)
		return nil, findErr
	}

	if qrcodeInfo != nil {
		qrcodeDelErr := l.svcCtx.ModelUserServiceQrcodeModel.Delete(l.ctx, qrcodeInfo.Id)
		if qrcodeDelErr != nil {
			l.Logger.Error("ModelUserServiceQrcodeModel_Update_Err", qrcodeDelErr)
			return nil, qrcodeDelErr
		}

		//结束语
		l.svcCtx.ModelUserServiceQrcodeConclusion.Delete(l.ctx, qrcodeInfo.Id)
	}

	return &wechat.ErrorResp{}, nil
}
