package model

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserTagModel = (*customTbExternalUserTagModel)(nil)

type (
	// TbExternalUserTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserTagModel.
	TbExternalUserTagModel interface {
		tbExternalUserTagModel
		FindListByExternalTagId(ctx context.Context, tagIdlist []string) ([]*TbExternalUserTag, error)
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

func (m *customTbExternalUserTagModel) FindListByExternalTagId(ctx context.Context, tagId []string) ([]*TbExternalUserTag, error) {

	var resp []*TbExternalUserTag
	if len(tagId) == 0 {
		return resp, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserRows).From(m.table).Where(squirrel.Eq{
		"tag_id": tagId,
	}).ToSql()
	if err != nil {
		return resp, err
	}

	err = m.conn.QueryRowsCtx(ctx, &resp, sql, args...)

	return resp, err
}
