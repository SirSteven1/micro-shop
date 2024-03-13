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
	userReceiveAddressFieldNames          = builder.RawFieldNames(&UserReceiveAddress{})
	userReceiveAddressRows                = strings.Join(userReceiveAddressFieldNames, ",")
	userReceiveAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(userReceiveAddressFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userReceiveAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(userReceiveAddressFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUserReceiveAddressIdPrefix = "cache:userReceiveAddress:id:"
)

type (
	userReceiveAddressModel interface {
		Insert(ctx context.Context, data *UserReceiveAddress) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserReceiveAddress, error)
		Update(ctx context.Context, data *UserReceiveAddress) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserReceiveAddressModel struct {
		sqlc.CachedConn
		table string
	}

	UserReceiveAddress struct {
		Id            int64     `db:"id"`
		Uid           uint64    `db:"uid"`            // 用户id
		Name          string    `db:"name"`           // 收货人名称
		Phone         string    `db:"phone"`          // 手机号
		IsDefault     uint64    `db:"is_default"`     // 是否为默认地址
		PostCode      string    `db:"post_code"`      // 邮政编码
		Province      string    `db:"province"`       // 省份/直辖市
		City          string    `db:"city"`           // 城市
		Region        string    `db:"region"`         // 区
		DetailAddress string    `db:"detail_address"` // 详细地址(街道)
		IsDelete      uint64    `db:"is_delete"`      // 是否删除
		CreateTime    time.Time `db:"create_time"`    // 数据创建时间[禁止在代码中赋值]
		UpdateTime    time.Time `db:"update_time"`    // 数据更新时间[禁止在代码中赋值]
	}
)

func newUserReceiveAddressModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserReceiveAddressModel {
	return &defaultUserReceiveAddressModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_receive_address`",
	}
}

func (m *defaultUserReceiveAddressModel) Delete(ctx context.Context, id int64) error {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userReceiveAddressIdKey)
	return err
}

func (m *defaultUserReceiveAddressModel) FindOne(ctx context.Context, id int64) (*UserReceiveAddress, error) {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, id)
	var resp UserReceiveAddress
	err := m.QueryRowCtx(ctx, &resp, userReceiveAddressIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userReceiveAddressRows, m.table)
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

func (m *defaultUserReceiveAddressModel) Insert(ctx context.Context, data *UserReceiveAddress) (sql.Result, error) {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userReceiveAddressRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Name, data.Phone, data.IsDefault, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.IsDelete)
	}, userReceiveAddressIdKey)
	return ret, err
}

func (m *defaultUserReceiveAddressModel) Update(ctx context.Context, data *UserReceiveAddress) error {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userReceiveAddressRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Name, data.Phone, data.IsDefault, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.IsDelete, data.Id)
	}, userReceiveAddressIdKey)
	return err
}

func (m *defaultUserReceiveAddressModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, primary)
}

func (m *defaultUserReceiveAddressModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userReceiveAddressRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserReceiveAddressModel) tableName() string {
	return m.table
}
