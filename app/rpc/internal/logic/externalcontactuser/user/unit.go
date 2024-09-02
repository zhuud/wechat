package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/wechat"
)

type GetExternalUserUnitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserUnitLogic {
	return &GetExternalUserUnitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type unitIds struct {
	RemarkTag []string
	Video     []string
}

func (t *GetExternalUserUnitLogic) GetUserUint(externalUser map[string]*types.ExternalUser, in *wechat.ExternalUserInfoReq) {
	//记录入参日志
	logc.Info(t.ctx, in, `params_unit`)

	ids := t.builds(externalUser)

	// 并行处理数据读取
	group := threading.NewRoutineGroup()

	//获取要查询的字段
	uintOpt := NewGetExternalUserCacheLogic(t.ctx, t.svcCtx).GetUField(in)

	for _, unitField := range uintOpt {
		group.RunSafe(func() {
			t.getUnit(ids, unitField, externalUser)
		})
	}
	group.Wait()

	return

}

func (t *GetExternalUserUnitLogic) builds(externalUserList map[string]*types.ExternalUser) (ids unitIds) {
	if externalUserList == nil || len(externalUserList) == 0 {
		return ids
	}

	for _, externalUser := range externalUserList {
		if externalUser.ExternalUserFollowAttribute == nil {
			continue
		}

		if externalUser.ExternalUserFollowAttribute.RemarkTag != nil {
			ids.RemarkTag = append(ids.RemarkTag, externalUser.ExternalUserFollowAttribute.RemarkTag...)
		}

		if externalUser.ExternalUserFollowAttribute.Video != nil {
			ids.Video = append(ids.RemarkTag, externalUser.ExternalUserFollowAttribute.Video...)
		}
	}

	return
}

func (t *GetExternalUserUnitLogic) getUnit(ids unitIds, unitField string, externalUser map[string]*types.ExternalUser) {

}
