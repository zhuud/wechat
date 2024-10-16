// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbExternalUserFieldNames          = builder.RawFieldNames(&TbExternalUser{})
	tbExternalUserRows                = strings.Join(tbExternalUserFieldNames, ",")
	tbExternalUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tbExternalUserFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbExternalUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tbExternalUserFieldNames, "`external_userid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbExternalUserModel interface {
		Insert(ctx context.Context, data *TbExternalUser) (sql.Result, error)
		FindOne(ctx context.Context, externalUserid string) (*TbExternalUser, error)
		Update(ctx context.Context, data *TbExternalUser) error
		Delete(ctx context.Context, externalUserid string) error
	}

	defaultTbExternalUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbExternalUser struct {
		ExternalUserid string    `db:"external_userid"` // 外部联系人的userid | 2020-09-10
		Unionid        string    `db:"unionid"`         // 外部联系人在微信开放平台的唯一身份标识（联系人类型是微信用户且企业绑定了微信开发者ID有此字段 第三方应用和代开发应用均不可获取 上游企业不可获取下游企业客户该字段） | 2020-09-10
		Type           int64     `db:"type"`            // 外部联系人的类型 (1:微信用户 / 2:企业微信用户) | 2020-09-10
		Name           string    `db:"name"`            // 外部联系人的名称(微信用户返回其微信昵称 企业微信联系人返回其设置对外展示的别名或实名) | 2020-09-10
		Avatar         string    `db:"avatar"`          // 外部联系人头像(代开发自建应用需要管理员授权才可以获取 第三方不可获取 上游企业不可获取下游企业客户该字段) | 2020-09-10
		Gender         int64     `db:"gender"`          // 外部联系人性别 (0:未知 / 1:男性 / 2:女性)(第三方不可获取 上游企业不可获取下游企业客户该字段 返回值为0) | 2020-09-10
		CorpName       string    `db:"corp_name"`       // 外部联系人所在企业的简称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10
		CorpFullName   string    `db:"corp_full_name"`  // 外部联系人所在企业的主体名称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10
		Position       string    `db:"position"`        // 外部联系人的职位(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10
		Status         int64     `db:"status"`          // 状态 (0:删除 / 1:正常) | 2020-09-10
		CreatedAt      time.Time `db:"created_at"`      // 创建时间 | 2020-09-10
		UpdatedAt      time.Time `db:"updated_at"`      // 更新时间 | 2020-09-10
	}
)

func newTbExternalUserModel(conn sqlx.SqlConn) *defaultTbExternalUserModel {
	return &defaultTbExternalUserModel{
		conn:  conn,
		table: "`tb_external_user`",
	}
}

func (m *defaultTbExternalUserModel) Delete(ctx context.Context, externalUserid string) error {
	query := fmt.Sprintf("delete from %s where `external_userid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, externalUserid)
	return err
}

func (m *defaultTbExternalUserModel) FindOne(ctx context.Context, externalUserid string) (*TbExternalUser, error) {
	query := fmt.Sprintf("select %s from %s where `external_userid` = ? limit 1", tbExternalUserRows, m.table)
	var resp TbExternalUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, externalUserid)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbExternalUserModel) Insert(ctx context.Context, data *TbExternalUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbExternalUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ExternalUserid, data.Unionid, data.Type, data.Name, data.Avatar, data.Gender, data.CorpName, data.CorpFullName, data.Position, data.Status)
	return ret, err
}

func (m *defaultTbExternalUserModel) Update(ctx context.Context, data *TbExternalUser) error {
	query := fmt.Sprintf("update %s set %s where `external_userid` = ?", m.table, tbExternalUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Unionid, data.Type, data.Name, data.Avatar, data.Gender, data.CorpName, data.CorpFullName, data.Position, data.Status, data.ExternalUserid)
	return err
}

func (m *defaultTbExternalUserModel) tableName() string {
	return m.table
}
