package externalcontactgroupchatlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalGroupChatInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalGroupChatInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalGroupChatInfoLogic {
	return &GetExternalGroupChatInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalGroupChatInfoLogic) GetExternalGroupChatInfo(in *wechat.ExternalGroupChatInfoReq) (*wechat.ErrorResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ErrorResp{}, nil
}
