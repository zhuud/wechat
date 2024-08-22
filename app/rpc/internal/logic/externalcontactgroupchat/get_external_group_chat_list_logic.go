package externalcontactgroupchatlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/wechat"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalGroupChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalGroupChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalGroupChatListLogic {
	return &GetExternalGroupChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalGroupChatListLogic) GetExternalGroupChatList(in *wechat.ExternalGroupChatListReq) (*wechat.ErrorResp, error) {
	// todo: add your logic here and delete this line

	return &wechat.ErrorResp{}, nil
}
