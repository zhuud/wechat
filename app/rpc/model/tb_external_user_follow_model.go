package model

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserFollowModel = (*customTbExternalUserFollowModel)(nil)

const (
	DelStatus            = 0 //互相删除
	NormalStatus         = 1 //正常
	FollowUserDelCStatus = 2 //客服删除用户
	CDelFollowUserStatus = 3 //用户删除客服
)

type (
	// TbExternalUserFollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserFollowModel.
	TbExternalUserFollowModel interface {
		tbExternalUserFollowModel
		FindListByExternalUserId(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollow, error)
		FindOneByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid, crop string) (*TbExternalUserFollow, error)
	}

	customTbExternalUserFollowModel struct {
		*defaultTbExternalUserFollowModel
	}
)

// NewTbExternalUserFollowModel returns a model for the database table.
func NewTbExternalUserFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbExternalUserFollowModel {
	return &customTbExternalUserFollowModel{
		defaultTbExternalUserFollowModel: newTbExternalUserFollowModel(conn, c, opts...),
	}
}

func (m *defaultTbExternalUserFollowModel) FindListByExternalUserId(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollow, error) {

	var resp []*TbExternalUserFollow
	if len(externalUserid) == 0 {
		return resp, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserFollowRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"status":          1,
	}).ToSql()
	if err != nil {
		return resp, err
	}

	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, args...)

	return resp, err
}

func (m *defaultTbExternalUserFollowModel) FindOneByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid, crop string) (*TbExternalUserFollow, error) {

	var resp TbExternalUserFollow
	if len(externalUserid) == 0 || len(userid) == 0 {
		return nil, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserFollowRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"userid":          userid,
		"platform":        crop,
		"status":          1,
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = m.QueryRowNoCacheCtx(ctx, &resp, sql, args...)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
