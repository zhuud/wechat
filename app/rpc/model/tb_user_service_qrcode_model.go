package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbUserServiceQrcodeModel = (*customTbUserServiceQrcodeModel)(nil)

type (
	// TbUserServiceQrcodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserServiceQrcodeModel.
	TbUserServiceQrcodeModel interface {
		tbUserServiceQrcodeModel
		FindOneByConfigId(ctx context.Context, configId string) (*TbUserServiceQrcode, error)
	}

	customTbUserServiceQrcodeModel struct {
		*defaultTbUserServiceQrcodeModel
	}
)

// NewTbUserServiceQrcodeModel returns a model for the database table.
func NewTbUserServiceQrcodeModel(conn sqlx.SqlConn) TbUserServiceQrcodeModel {
	return &customTbUserServiceQrcodeModel{
		defaultTbUserServiceQrcodeModel: newTbUserServiceQrcodeModel(conn),
	}
}

func (m *customTbUserServiceQrcodeModel) FindOneByConfigId(ctx context.Context, configId string) (*TbUserServiceQrcode, error) {
	query := fmt.Sprintf("select %s from %s where `config_id` = ? limit 1", tbUserServiceQrcodeRows, m.table)
	var resp TbUserServiceQrcode
	err := m.conn.QueryRowCtx(ctx, &resp, query, configId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
