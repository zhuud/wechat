package model

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbExternalUserAttributeModel = (*customTbExternalUserAttributeModel)(nil)

const (
	TbExternalUserAttrNormalStatus = 1
	AttributeTypeText              = 0
	AttributeTypeWeb               = 1
	AttributeTypeMiniprogram       = 2
	AttributeTypeprofile           = 3
)

type (
	// TbExternalUserAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserAttributeModel.
	TbExternalUserAttributeModel interface {
		tbExternalUserAttributeModel
		FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserAttribute, error)
		DeleteByExternalUserId(ctx context.Context, externalUserid string) error
	}

	customTbExternalUserAttributeModel struct {
		*defaultTbExternalUserAttributeModel
	}
)

// NewTbExternalUserAttributeModel returns a model for the database table.
func NewTbExternalUserAttributeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbExternalUserAttributeModel {
	return &customTbExternalUserAttributeModel{
		defaultTbExternalUserAttributeModel: newTbExternalUserAttributeModel(conn, c, opts...),
	}
}

func (m *defaultTbExternalUserAttributeModel) DeleteByExternalUserId(ctx context.Context, externalUserid string) error {

	if len(externalUserid) == 0 {
		return errors.New("参数为空")
	}
	sql, args, err := squirrel.Update("status").From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"status":          1,
	}).ToSql()

	if err != nil {
		return err
	}
	_, err = m.ExecNoCacheCtx(ctx, sql, args...)

	return err
}

func (m *defaultTbExternalUserAttributeModel) FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserAttribute, error) {

	var resp []*TbExternalUserAttribute
	if len(externalUserid) == 0 {
		return resp, errors.New("参数为空")
	}

	sql, args, err := squirrel.Select(tbExternalUserAttributeRows).From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
	}).ToSql()
	if err != nil {
		return resp, err
	}

	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, args...)

	return resp, err
}
