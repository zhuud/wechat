package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbUserServiceQrcodeConclusionsModel = (*customTbUserServiceQrcodeConclusionsModel)(nil)

type (
	// TbUserServiceQrcodeConclusionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserServiceQrcodeConclusionsModel.
	TbUserServiceQrcodeConclusionsModel interface {
		tbUserServiceQrcodeConclusionsModel
	}

	customTbUserServiceQrcodeConclusionsModel struct {
		*defaultTbUserServiceQrcodeConclusionsModel
	}
)

// NewTbUserServiceQrcodeConclusionsModel returns a model for the database table.
func NewTbUserServiceQrcodeConclusionsModel(conn sqlx.SqlConn) TbUserServiceQrcodeConclusionsModel {
	return &customTbUserServiceQrcodeConclusionsModel{
		defaultTbUserServiceQrcodeConclusionsModel: newTbUserServiceQrcodeConclusionsModel(conn),
	}
}
