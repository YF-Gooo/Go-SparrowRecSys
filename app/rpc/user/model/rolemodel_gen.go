// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roleFieldNames          = builder.RawFieldNames(&Role{})
	roleRows                = strings.Join(roleFieldNames, ",")
	roleRowsExpectAutoSet   = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	roleRowsWithPlaceHolder = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheRoleIdPrefix = "cache:role:id:"
)

type (
	roleModel interface {
		Insert(ctx context.Context, data *Role) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Role, error)
		Update(ctx context.Context, data *Role) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRoleModel struct {
		sqlc.CachedConn
		table string
	}

	Role struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		Status     int64     `db:"status"`
		RoleName   string    `db:"role_name"`
		RoleDesc   string    `db:"role_desc"`
	}
)

func newRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRoleModel {
	return &defaultRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`role`",
	}
}

func (m *defaultRoleModel) Delete(ctx context.Context, id int64) error {
	roleIdKey := fmt.Sprintf("%s%v", cacheRoleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, roleIdKey)
	return err
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id int64) (*Role, error) {
	roleIdKey := fmt.Sprintf("%s%v", cacheRoleIdPrefix, id)
	var resp Role
	err := m.QueryRowCtx(ctx, &resp, roleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	roleIdKey := fmt.Sprintf("%s%v", cacheRoleIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, roleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.Status, data.RoleName, data.RoleDesc)
	}, roleIdKey)
	return ret, err
}

func (m *defaultRoleModel) Update(ctx context.Context, data *Role) error {
	roleIdKey := fmt.Sprintf("%s%v", cacheRoleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.Status, data.RoleName, data.RoleDesc, data.Id)
	}, roleIdKey)
	return err
}

func (m *defaultRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRoleIdPrefix, primary)
}

func (m *defaultRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRoleModel) tableName() string {
	return m.table
}
