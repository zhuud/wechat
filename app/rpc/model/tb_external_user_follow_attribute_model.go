package model

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserFollowAttributeModel = (*customTbExternalUserFollowAttributeModel)(nil)

const (
	AttributeTypeRemarkTag = 1
	AttributeTypeVideo     = 2
)

type (
	// TbExternalUserFollowAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserFollowAttributeModel.
	TbExternalUserFollowAttributeModel interface {
		tbExternalUserFollowAttributeModel
		FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollowAttribute, error)
		DeleteByExternalUserIdAndUserIdAndPlatform(ctx context.Context, externalUserid, userid, crop string) error
	}

	customTbExternalUserFollowAttributeModel struct {
		*defaultTbExternalUserFollowAttributeModel
	}
)

// NewTbExternalUserFollowAttributeModel returns a model for the database table.
func NewTbExternalUserFollowAttributeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbExternalUserFollowAttributeModel {
	return &customTbExternalUserFollowAttributeModel{
		defaultTbExternalUserFollowAttributeModel: newTbExternalUserFollowAttributeModel(conn, c, opts...),
	}
}

func (m *defaultTbExternalUserFollowAttributeModel) DeleteByExternalUserIdAndUserIdAndPlatform(ctx context.Context, externalUserid, userid, crop string) error {

	if len(externalUserid) == 0 || len(userid) == 0 {
		return errors.New("参数为空")
	}
	sql, args, err := squirrel.Update("status").From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"userid":          externalUserid,
		"platform":        crop,
		"status":          1,
	}).ToSql()

	if err != nil {
		return err
	}
	_, err = m.ExecNoCacheCtx(ctx, sql, args...)

	return err
}

func (m *defaultTbExternalUserFollowAttributeModel) FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollowAttribute, error) {

	var resp []*TbExternalUserFollowAttribute
	if len(externalUserid) == 0 {
		return resp, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserFollowAttributeRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
	}).ToSql()
	if err != nil {
		return resp, err
	}

	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, args...)

	return resp, err
}
