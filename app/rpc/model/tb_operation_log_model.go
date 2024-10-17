package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbOperationLogModel = (*customTbOperationLogModel)(nil)

type (
	// TbOperationLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbOperationLogModel.
	TbOperationLogModel interface {
		tbOperationLogModel
	}

	customTbOperationLogModel struct {
		*defaultTbOperationLogModel
	}
)

// NewTbOperationLogModel returns a model for the database table.
func NewTbOperationLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbOperationLogModel {
	return &customTbOperationLogModel{
		defaultTbOperationLogModel: newTbOperationLogModel(conn, c, opts...),
	}
}
