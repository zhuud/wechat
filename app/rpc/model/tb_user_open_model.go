package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserOpenModel = (*customTbUserOpenModel)(nil)

type (
	// TbUserOpenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserOpenModel.
	TbUserOpenModel interface {
		tbUserOpenModel
	}

	customTbUserOpenModel struct {
		*defaultTbUserOpenModel
	}
)

// NewTbUserOpenModel returns a model for the database table.
func NewTbUserOpenModel(conn sqlx.SqlConn) TbUserOpenModel {
	return &customTbUserOpenModel{
		defaultTbUserOpenModel: newTbUserOpenModel(conn),
	}
}
