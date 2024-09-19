package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/model"
)

type GetExternalUserTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserTagLogic {
	return &GetExternalUserTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 获取用户标签信息
 */
func (t *GetExternalUserTagLogic) GetUserTagByTagIdList(tagIdList []string) (externalUserMap map[string]*model.TbExternalUserTag, err error) {
	externalUserTagList, err := t.svcCtx.ModelExternalUserTag.FindListByExternalTagId(t.ctx, tagIdList)
	if err != nil {
		logc.Error(t.ctx, `GetUserFollowAttributeByExternalUserIdList_err`, err)
		return nil, err
	}

	for _, externalUserTag := range externalUserTagList {
		externalUserMap[externalUserTag.TagId] = externalUserTag
	}

	return
}
