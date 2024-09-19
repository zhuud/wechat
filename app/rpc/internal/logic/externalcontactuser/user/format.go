package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
	"rpc/wechat"
)

type GetExternalUserFormatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalUserFormatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserFormatLogic {
	return &GetExternalUserFormatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (t *GetExternalUserFormatLogic) HandleFollowUserFormat(externalUser *types.ExternalUser, unit UserUnit) {
	//标签拆分
	tagMap := t.HandleFollowUserTagFormat(unit.Tag, externalUser.ExternalUserFollowAttributeDB)
	for k, followInfo := range externalUser.ExternalUserFollow {
		//视频号信息
		if video, ok := unit.Video[followInfo.Userid]; ok && followInfo.AddWay == 10 {
			externalUser.ExternalUserFollow[k].WechatChannels = video
		}

		//标签信息
		if tagInfo, ok := tagMap[followInfo.Userid]; ok {
			externalUser.ExternalUserFollow[k].Tags = tagInfo
		}

	}
}

func (t *GetExternalUserFormatLogic) HandleFollowUserTagFormat(tagList map[string]*model.TbExternalUserTag, attributeList map[string]*model.TbExternalUserFollowAttribute) map[string][]*wechat.ExternalUserFollowUserTag {
	tagUserList := map[string][]*wechat.ExternalUserFollowUserTag{}
	for _, attribute := range attributeList {
		if attribute.AttributeType != model.AttributeTypeRemarkTag {
			continue
		}

		if tagInfo, ok := tagList[attribute.AttributeValue]; ok {
			tagUser := &wechat.ExternalUserFollowUserTag{
				GroupName: tagInfo.GroupName,
				TagName:   tagInfo.Name,
				TagId:     tagInfo.TagId,
				Type:      0, //todo 标签类型
			}

			tagUserList[attribute.Userid] = append(tagUserList[attribute.Userid], tagUser)
		}
	}

	return tagUserList
}
