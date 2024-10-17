package model

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	TbExternalUserFollowAttrNormalStatus = 1
	AttributeTypeRemarkTag               = 1
	AttributeTypeVideo                   = 2
)

var _ TbExternalUserFollowAttributeModel = (*customTbExternalUserFollowAttributeModel)(nil)

type (
	// TbExternalUserFollowAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbExternalUserFollowAttributeModel.
	TbExternalUserFollowAttributeModel interface {
		tbExternalUserFollowAttributeModel
		FindListByExternalUserid(ctx context.Context, externalUserid []string) ([]*TbExternalUserFollowAttribute, error)
		DeleteByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid, crop string) error
	}

	customTbExternalUserFollowAttributeModel struct {
		*defaultTbExternalUserFollowAttributeModel
	}
)

// NewTbExternalUserFollowAttributeModel returns a model for the database table.
func NewTbExternalUserFollowAttributeModel(conn sqlx.SqlConn) TbExternalUserFollowAttributeModel {
	return &customTbExternalUserFollowAttributeModel{
		defaultTbExternalUserFollowAttributeModel: newTbExternalUserFollowAttributeModel(conn),
	}
}

func (m *defaultTbExternalUserFollowAttributeModel) DeleteByExternalUserIdAndUserId(ctx context.Context, externalUserid, userid, crop string) error {

	if len(externalUserid) == 0 || len(userid) == 0 {
		return errors.New("参数为空")
	}
	sql, args, err := squirrel.Update("status").From(m.table).Where(squirrel.Eq{
		"external_userid": externalUserid,
		"userid":          externalUserid,
		"crop":            crop,
		"status":          TbExternalUserFollowAttrNormalStatus,
	}).ToSql()

	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, sql, args...)

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

	err = m.conn.QueryRowsCtx(ctx, &resp, sql, args...)

	return resp, err
}
