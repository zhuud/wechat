package externalcontactuserlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveExternalUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExternalUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExternalUserFollowLogic {
	return &SaveExternalUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 删除关系
 */
func (l *SaveExternalUserFollowLogic) DeleteExternalUserFollow(crop string, externalUser types.ExternalData, status int) error {
	if externalUser.ExternalUserID == `` || externalUser.UserID == `` {
		return nil
	}

	update := model.TbExternalUserFollow{
		ExternalUserid: externalUser.ExternalUserID,
		Userid:         externalUser.UserID,
		Status:         cast.ToInt64(status),
		DeletedAt:      externalUser.CreateTime,
		Crop:           crop,
	}
	err := l.svcCtx.ModelExternalUserFollow.Update(l.ctx, &update)
	return err
}

/**
 * 编辑同意状态
 */
func (l *SaveExternalUserFollowLogic) UpdateChatAgreeStatus(externalUserId string, userId string) error {
	if externalUserId == `` || userId == `` {
		return nil
	}

	update := model.TbExternalUserFollow{
		ExternalUserid:  externalUserId,
		Userid:          userId,
		ChatAgreeStatus: 1,
	}
	err := l.svcCtx.ModelExternalUserFollow.Update(l.ctx, &update)

	return err
}

/**
 * 给用户添加的所有员工添加描述
 */
func (l *SaveExternalUserFollowLogic) UpdateFollowRemark() {

}
