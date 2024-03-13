package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecommendModel = (*customRecommendModel)(nil)

type (
	// RecommendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecommendModel.
	RecommendModel interface {
		recommendModel
	}

	customRecommendModel struct {
		*defaultRecommendModel
	}
)

// NewRecommendModel returns a model for the database table.
func NewRecommendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RecommendModel {
	return &customRecommendModel{
		defaultRecommendModel: newRecommendModel(conn, c, opts...),
	}
}
