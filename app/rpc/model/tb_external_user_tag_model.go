package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbExternalUserTagModel = (*customTbExternalUserTagModel)(nil)

type (
	// TbExternalUserTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserTagModel.
	TbExternalUserTagModel interface {
		tbExternalUserTagModel
	}

	customTbExternalUserTagModel struct {
		*defaultTbExternalUserTagModel
	}
)

// NewTbExternalUserTagModel returns a model for the database table.
func NewTbExternalUserTagModel(conn sqlx.SqlConn) TbExternalUserTagModel {
	return &customTbExternalUserTagModel{
		defaultTbExternalUserTagModel: newTbExternalUserTagModel(conn),
	}
}
