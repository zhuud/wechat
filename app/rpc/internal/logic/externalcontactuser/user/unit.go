package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
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
}

type UserUnit struct {
	Tag   map[string]*model.TbExternalUserTag
	Video map[string]*wechat.ExternalUserFollowUserWechatChannel
}

func (t *GetExternalUserUnitLogic) GetUserUint(externalUser map[string]*types.ExternalUser, in *wechat.ExternalUserInfoReq) (userUnit UserUnit) {
	//记录入参日志
	logc.Info(t.ctx, in, `params_unit`)

	ids := t.builds(externalUser)

	//需要查询的数据
	uintOpt := t.getUField(ids)

	// 并行处理数据读取
	group := threading.NewRoutineGroup()

	for _, unitField := range uintOpt {
		group.RunSafe(func() {
			t.getUnit(ids, unitField, externalUser, &userUnit)
		})
	}
	group.Wait()

	return userUnit

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
	}

	return
}

func (t *GetExternalUserUnitLogic) getUField(ids unitIds) (opt []string) {
	if ids.RemarkTag != nil && len(ids.RemarkTag) > 0 {
		opt = append(opt, `follow_user_tag`)
	}

	return
}

func (t *GetExternalUserUnitLogic) getUnit(ids unitIds, unitField string, externalUser map[string]*types.ExternalUser, userUnit *UserUnit) {
	switch unitField {
	case `follow_user_tag`:
		if ids.RemarkTag == nil || len(ids.RemarkTag) == 0 {
			return
		}
		tagList, err := NewGetExternalUserTagLogic(t.ctx, t.svcCtx).GetUserTagByTagIdList(ids.RemarkTag)
		if err != nil {
			t.Logger.Error(t.ctx, `GetUserTagByTagIdList_err`, err)
		}

		userUnit.Tag = tagList
	}

	return
}
