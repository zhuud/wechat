package model

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
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
		FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUser, error)
	}

	customTbExternalUserModel struct {
		*defaultTbExternalUserModel
	}
)

// NewTbExternalUserModel returns a model for the database table.
func NewTbExternalUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbExternalUserModel {
	return &customTbExternalUserModel{
		defaultTbExternalUserModel: newTbExternalUserModel(conn, c, opts...),
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

	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, args...)

	return resp, err
}
