package model

import (
    "context"
    "errors"
    "github.com/Masterminds/squirrel"
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserFollowModel = (*customTbExternalUserFollowModel)(nil)

type (
    // TbExternalUserFollowModel is an interface to be customized, add more methods here,
    // and implement the added methods in customTbExternalUserFollowModel.
    TbExternalUserFollowModel interface {
        tbExternalUserFollowModel
        FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollow, error)
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

func (m *defaultTbExternalUserFollowModel) FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollow, error) {

    var resp []*TbExternalUserFollow
    if len(externalUserid) == 0 {
        return resp, errors.New("参数为空")
    }

    sql, args, err := squirrel.Select(tbExternalUserFollowRows).From(m.table).Where(squirrel.Eq{
        "external_userid": externalUserid,
    }).ToSql()
    if err != nil {
        return resp, err
    }

    err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, args...)

    return resp, err
}
