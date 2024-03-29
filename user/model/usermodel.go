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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserIdPrefix        = "cache:user:id:"
	cacheUserUserNamePrefix  = "cache:user:userName:"
	cacheUserUserPhonePrefix = "cache:user:userPhone:"
	cacheUserIdentityPrefix  = "cache:user:userIdentity:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*User, error)
		FindOneByUserName(ctx context.Context, userName string) (*User, error)
		FindOneByUserPhone(ctx context.Context, userPhone string) (*User, error)
		FindOneByUserIdentity(ctx context.Context, UserIdentity string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id           int64     `db:"id"`           // 用户ID
		UserIdentity string    `db:"userIdentity"` // 用户唯一标识
		UserName     string    `db:"userName"`     // 用户名
		PassWord     string    `db:"passWord"`     // 用户密码，MD5加密
		UserNick     string    `db:"userNick"`     // 用户昵称
		UserFace     string    `db:"userFace"`     // 用户头像地址
		UserSex      int64     `db:"UserSex"`      // 用户性别：0男，1女，2保密
		UserEmail    string    `db:"userEmail"`    // 用户邮箱
		UserPhone    string    `db:"userPhone"`    // 手机号
		LoginAddress string    `db:"loginAddress"` // 用户登录IP地址
		CreateTime   time.Time `db:"create_time"`  // 创建时间
		UpdateTime   time.Time `db:"update_time"`  // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, data.UserName)
	userUserPhoneKey := fmt.Sprintf("%s%v", cacheUserUserPhonePrefix, data.UserPhone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey, userUserNameKey, userUserPhoneKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id uint64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByUserName(ctx context.Context, userName string) (*User, error) {
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, userName)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `userName` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userName); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUserPhone(ctx context.Context, userPhone string) (*User, error) {
	userUserPhoneKey := fmt.Sprintf("%s%v", cacheUserUserPhonePrefix, userPhone)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userUserPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `userPhone` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userPhone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUserIdentity(ctx context.Context, userIdentity string) (*User, error) {
	userUserIdentityKey := fmt.Sprintf("%s%v", cacheUserIdentityPrefix, userIdentity)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userUserIdentityKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `userIdentity` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userIdentity); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, data.UserName)
	userUserPhoneKey := fmt.Sprintf("%s%v", cacheUserUserPhonePrefix, data.UserPhone)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserIdentity, data.UserName, data.PassWord, data.UserNick, data.UserFace, data.UserSex, data.UserEmail, data.UserPhone, data.LoginAddress)
	}, userIdKey, userUserNameKey, userUserPhoneKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, uint64(newData.Id))
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, data.UserName)
	userUserPhoneKey := fmt.Sprintf("%s%v", cacheUserUserPhonePrefix, data.UserPhone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserIdentity, newData.UserName, newData.PassWord, newData.UserNick, newData.UserFace, newData.UserSex, newData.UserEmail, newData.UserPhone, newData.LoginAddress, newData.Id)
	}, userIdKey, userUserNameKey, userUserPhoneKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
