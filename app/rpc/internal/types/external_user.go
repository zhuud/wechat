package types

import (
	"rpc/model"
	"rpc/wechat"
)

type ExternalUser struct {
	ExternalUser                  *wechat.ExternalUser
	ExternalUserFollow            []*wechat.ExternalUserFollowUser
	ExternalUserFollowAttribute   *ExternalUserFollowAttribute
	ExternalUserFollowAttributeDB map[string]*model.TbExternalUserFollowAttribute
}

type ExternalUserUnit struct {
	ExternalUser                map[string]*wechat.ExternalUser
	ExternalUserAttribute       map[string][]*model.TbExternalUserAttribute
	ExternalUserFollow          map[string][]*wechat.ExternalUserFollowUser
	ExternalUserFollowAttribute map[string]map[string]*model.TbExternalUserFollowAttribute
}

type ExternalUserFollowAttribute struct {
	RemarkTag []string
	Video     map[string]*wechat.ExternalUserFollowUserWechatChannel
}
