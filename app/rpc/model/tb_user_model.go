package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbUserModel = (*customTbUserModel)(nil)

type (
	// TbUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserModel.
	TbUserModel interface {
		tbUserModel
	}

	customTbUserModel struct {
		*defaultTbUserModel
	}
)

// NewTbUserModel returns a model for the database table.
func NewTbUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbUserModel {
	return &customTbUserModel{
		defaultTbUserModel: newTbUserModel(conn, c, opts...),
	}
}
