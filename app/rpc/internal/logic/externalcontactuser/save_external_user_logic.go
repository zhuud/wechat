package externalcontactuserlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
)

type SaveExternalUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExternalUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExternalUserLogic {
	return &SaveExternalUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveExternalUserLogic) SyncExternalUser(externalUserData types.ExternalData) {
	if externalUserData.ExternalUserID == `` || externalUserData.ToUserName == `` {
		return
	}
	externalUser, err := l.svcCtx.WeCom.WithCorp(`yx`).ExternalUser.Get(l.ctx, externalUserData.ExternalUserID, `CURSOR`)
	if err != nil || externalUser == nil {
		return
	}

}

/**
 * 删除关系
 */
func (l *SaveExternalUserLogic) DeleteExternalUserFollow(externalUserId string, userId string, dataOrigin string) {
	if externalUserId == `` || userId == `` {
		return
	}

	status := 0
	if dataOrigin == `del_external_contact` {
		status = model.ExternalUserDelStatus
	} else if dataOrigin == `del_follow_user` {
		status = model.StaffDelStatus
	}

	update := model.TbExternalUserFollow{
		ExternalUserid: externalUserId,
		Userid:         userId,
		Status:         status,
	}
	l.svcCtx.ModelExternalUserFollow.Update(l.ctx, &update)
}
