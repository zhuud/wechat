package user

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/model"
	"rpc/wechat"
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
 * 获取用户扩展信息
 */
func (t *GetExternalUserAttributeLogic) GetUserAttributeByExternalUserIdList(externalUserIdList []string) (externalUserMap map[string][]*model.TbExternalUserAttribute, err error) {
	externalUserAttributeList, err := t.svcCtx.ModelExternalUserAttribute.FindListByExternalUserid(t.ctx, externalUserIdList)
	if err != nil {
		logc.Error(t.ctx, `GetUserFollowAttributeByExternalUserIdList_err`, err)
		return nil, err
	}

	for _, externalUserAttribute := range externalUserAttributeList {
		if externalUserMap[externalUserAttribute.ExternalUserid] == nil {
			externalUserMap[externalUserAttribute.ExternalUserid] = []*model.TbExternalUserAttribute{}
		}
		externalUserMap[externalUserAttribute.ExternalUserid] = append(externalUserMap[externalUserAttribute.ExternalUserid], externalUserAttribute)
	}

	return
}

func (t *GetExternalUserAttributeLogic) HandleUserExternalProfileAttribute(externalUserAttributeList map[string][]*model.TbExternalUserAttribute) (externalProfileMap map[string]wechat.ExternalUserProfile) {
	if externalUserAttributeList == nil || len(externalUserAttributeList) == 0 {
		return
	}

	for externalUserId, externalUserAttribute := range externalUserAttributeList {
		if externalUserAttribute == nil {
			continue
		}

		externalUserProfile := wechat.ExternalUserProfile{}
		externalAttr := []wechat.ExternalUserProfileItem{}
		for _, externalUserAttributeInfo := range externalUserAttribute {
			//文本
			if externalUserAttributeInfo.AttributeType == model.AttributeTypeText {
				text := wechat.ExternalUserProfileItemText{}
				json.Unmarshal([]byte(externalUserAttributeInfo.Extension), &text)
				externalAttr = append(externalAttr, wechat.ExternalUserProfileItem{
					Type: cast.ToInt32(externalUserAttributeInfo.AttributeType),
					Name: externalUserAttributeInfo.AttributeValue,
					Text: &text,
				})

				//网页
			} else if externalUserAttributeInfo.AttributeType == model.AttributeTypeWeb {
				//小程序
			} else if externalUserAttributeInfo.AttributeType == model.AttributeTypeMiniprogram {

			} else if externalUserAttributeInfo.AttributeType == model.AttributeTypeprofile {
			}
		}

		externalProfileMap[externalUserId] = externalUserProfile
	}

	return
}
