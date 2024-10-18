package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbUserOpenModel = (*customTbUserOpenModel)(nil)

type (
	// TbUserOpenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserOpenModel.
	TbUserOpenModel interface {
		tbUserOpenModel
		FindUidByUnionid(ctx context.Context, unionid string) int64
		FindUnionidByUid(ctx context.Context, uid int64) string
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

func (m *defaultTbUserOpenModel) FindUidByUnionid(ctx context.Context, unionid string) int64 {

	var resp TbUserOpen
	if len(unionid) == 0 {
		return 0
	}

	sql, args, err := squirrel.Select(tbUserOpenRows).From(m.table).Where(squirrel.Eq{
		"open_id":   unionid,
		"open_site": 3,
	}).ToSql()
	if err != nil {
		return 0
	}

	err = m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	if err != nil {
		logx.Errorf("FindUidByUnionid error : %v", err)
	}

	return resp.Uid
}

func (m *defaultTbUserOpenModel) FindUnionidByUid(ctx context.Context, uid int64) string {

	var resp TbUserOpen
	if uid == 0 {
		return ""
	}

	sql, args, err := squirrel.Select(tbUserOpenRows).From(m.table).Where(squirrel.Eq{
		"uid":       uid,
		"open_site": 3,
	}).ToSql()
	if err != nil {
		return ""
	}

	err = m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	if err != nil {
		logx.Errorf("FindUnionidByUid error : %v", err)
	}

	return resp.OpenId
}
