package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
)

type GetExternalUserAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserAttributeLogic {
	return &GetExternalUserAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 获取用户基础信息
 */
func (t *GetExternalUserAttributeLogic) GetUserFollowAttributeByExternalUserIdList(externalUserIdList []string) (externalUserMap map[string][]*model.TbExternalUserFollowAttribute, err error) {
	externalUserAttributeList, err := t.svcCtx.ModelExternalUserFollowAttribute.FindListByExternalUserid(t.ctx, externalUserIdList)
	if err != nil {
		logc.Error(t.ctx, `GetUserFollowAttributeByExternalUserIdList_err`, err)
		return nil, err
	}

	for _, externalUserAttribute := range externalUserAttributeList {
		externalUserMap[externalUserAttribute.ExternalUserid] = append(externalUserMap[externalUserAttribute.ExternalUserid], externalUserAttribute)
	}

	return
}

func (t *GetExternalUserAttributeLogic) HandleAttributeFormat(externalUserFollowAttrList []*model.TbExternalUserFollowAttribute) (externalUserAttribute *types.ExternalUserAttribute) {
	for _, externalUserFollowAttr := range externalUserFollowAttrList {
		//标签信息
		if externalUserFollowAttr.AttributeType == model.AttributeTypeRemarkTag {
			externalUserAttribute.RemarkTag = append(externalUserAttribute.RemarkTag, externalUserFollowAttr.AttributeValue)
		}

		//视频信息
		if externalUserFollowAttr.AttributeType == model.AttributeTypeVideo {
			externalUserAttribute.Video = append(externalUserAttribute.Video, externalUserFollowAttr.AttributeValue)
		}
	}

	return
}
