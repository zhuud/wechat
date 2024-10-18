package model

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbPrivateDomainUserModel = (*customTbPrivateDomainUserModel)(nil)

type (
	// TbPrivateDomainUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbPrivateDomainUserModel.
	TbPrivateDomainUserModel interface {
		tbPrivateDomainUserModel
		FindOneByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid string, from int) (*TbPrivateDomainUser, error)
	}

	customTbPrivateDomainUserModel struct {
		*defaultTbPrivateDomainUserModel
	}
)

// NewTbPrivateDomainUserModel returns a model for the database table.
func NewTbPrivateDomainUserModel(conn sqlx.SqlConn) TbPrivateDomainUserModel {
	return &customTbPrivateDomainUserModel{
		defaultTbPrivateDomainUserModel: newTbPrivateDomainUserModel(conn),
	}
}

func (m *defaultTbPrivateDomainUserModel) FindOneByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid string, from int) (*TbPrivateDomainUser, error) {

	var resp TbPrivateDomainUser
	if len(externalUserid) == 0 || len(userid) == 0 {
		return nil, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbPrivateDomainUserRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"qywx_user_id":    userid,
		"from":            from,
	}).OrderBy("id desc").Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	err = m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
