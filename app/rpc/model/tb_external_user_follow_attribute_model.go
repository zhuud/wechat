package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserFollowAttributeModel = (*customTbExternalUserFollowAttributeModel)(nil)

type (
	// TbExternalUserFollowAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserFollowAttributeModel.
	TbExternalUserFollowAttributeModel interface {
		tbExternalUserFollowAttributeModel
	}

	customTbExternalUserFollowAttributeModel struct {
		*defaultTbExternalUserFollowAttributeModel
	}
)

// NewTbExternalUserFollowAttributeModel returns a model for the database table.
func NewTbExternalUserFollowAttributeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbExternalUserFollowAttributeModel {
	return &customTbExternalUserFollowAttributeModel{
		defaultTbExternalUserFollowAttributeModel: newTbExternalUserFollowAttributeModel(conn, c, opts...),
	}
}
