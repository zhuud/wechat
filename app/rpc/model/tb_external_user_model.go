package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	TbExternalUserNormalStatus = 1
)

var _ TbExternalUserModel = (*customTbExternalUserModel)(nil)

type (
	// TbExternalUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserModel.
	TbExternalUserModel interface {
		tbExternalUserModel
		FindOne(ctx context.Context, externalUserid string) (*TbExternalUser, error)
		FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUser, error)
	}

	customTbExternalUserModel struct {
		*defaultTbExternalUserModel
	}
)

// NewTbExternalUserModel returns a model for the database table.
func NewTbExternalUserModel(conn sqlx.SqlConn) TbExternalUserModel {
	return &customTbExternalUserModel{
		defaultTbExternalUserModel: newTbExternalUserModel(conn),
	}
}

func (m *customTbExternalUserModel) FindOne(ctx context.Context, externalUserid string) (*TbExternalUser, error) {
	query := fmt.Sprintf("select %s from %s where `external_userid` = ? and status = %d limit 1", tbExternalUserRows, m.table, TbExternalUserNormalStatus)
	var resp TbExternalUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, externalUserid)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customTbExternalUserModel) FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUser, error) {

	var resp []*TbExternalUser
	if len(externalUserid) == 0 {
		return resp, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
	}).ToSql()
	if err != nil {
		return resp, err
	}

	err = m.conn.QueryRowsCtx(ctx, &resp, sql, args...)

	return resp, err
}
