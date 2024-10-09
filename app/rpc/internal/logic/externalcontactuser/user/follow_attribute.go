package user

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
	"rpc/wechat"
)

type GetExternalUserFollowAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserFollowAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserFollowAttributeLogic {
	return &GetExternalUserFollowAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 获取用户基础信息
 */
func (t *GetExternalUserFollowAttributeLogic) GetUserFollowAttributeByExternalUserIdList(externalUserIdList []string) (externalUserMap map[string]map[string]*model.TbExternalUserFollowAttribute, err error) {
	externalUserAttributeList, err := t.svcCtx.ModelExternalUserFollowAttribute.FindListByExternalUserid(t.ctx, externalUserIdList)
	if err != nil {
		logc.Error(t.ctx, `GetUserFollowAttributeByExternalUserIdList_err`, err)
		return nil, err
	}

	for _, externalUserAttribute := range externalUserAttributeList {
		externalUserMap[externalUserAttribute.ExternalUserid][externalUserAttribute.Userid] = externalUserAttribute
	}

	return
}

func (t *GetExternalUserFollowAttributeLogic) HandleAttributeFormat(externalUserFollowAttrList map[string]*model.TbExternalUserFollowAttribute) (externalUserAttribute *types.ExternalUserFollowAttribute) {
	if externalUserFollowAttrList == nil || len(externalUserFollowAttrList) == 0 {
		return
	}
	for _, externalUserFollowAttr := range externalUserFollowAttrList {
		//标签信息
		if externalUserFollowAttr.AttributeType == model.AttributeTypeRemarkTag {
			externalUserAttribute.RemarkTag = append(externalUserAttribute.RemarkTag, externalUserFollowAttr.AttributeValue)
		}

		//视频信息
		if externalUserFollowAttr.AttributeType == model.AttributeTypeVideo {
			video := &wechat.ExternalUserFollowUserWechatChannel{}
			json.Unmarshal([]byte(externalUserFollowAttr.Extension), &video)
			externalUserAttribute.Video[externalUserFollowAttr.Userid] = video
		}
	}

	return
}
