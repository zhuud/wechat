package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserServiceQrcodeModel = (*customUserServiceQrcodeModel)(nil)

type (
	// UserServiceQrcodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserServiceQrcodeModel.
	UserServiceQrcodeModel interface {
		userServiceQrcodeModel
		withSession(session sqlx.Session) UserServiceQrcodeModel
	}

	customUserServiceQrcodeModel struct {
		*defaultUserServiceQrcodeModel
	}
)

// NewUserServiceQrcodeModel returns a model for the database table.
func NewUserServiceQrcodeModel(conn sqlx.SqlConn) UserServiceQrcodeModel {
	return &customUserServiceQrcodeModel{
		defaultUserServiceQrcodeModel: newUserServiceQrcodeModel(conn),
	}
}

func (m *customUserServiceQrcodeModel) withSession(session sqlx.Session) UserServiceQrcodeModel {
	return NewUserServiceQrcodeModel(sqlx.NewSqlConnFromSession(session))
}
