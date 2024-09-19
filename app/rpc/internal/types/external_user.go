package types

import (
	"rpc/model"
	"rpc/wechat"
)

type ExternalUser struct {
	ExternalUser                  *wechat.ExternalUser
	ExternalUserFollow            []*wechat.ExternalUserFollowUser
	ExternalUserFollowAttribute   *ExternalUserAttribute
	ExternalUserFollowAttributeDB map[string]*model.TbExternalUserFollowAttribute
}

type ExternalUserUnit struct {
	ExternalUser                map[string]*wechat.ExternalUser
	ExternalUserFollow          map[string][]*wechat.ExternalUserFollowUser
	ExternalUserFollowAttribute map[string]map[string]*model.TbExternalUserFollowAttribute
}

type ExternalUserAttribute struct {
	RemarkTag []string
	Video     map[string]*wechat.ExternalUserFollowUserWechatChannel
}
