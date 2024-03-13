package svc

import (
	"fmt"
	"microShop/product/api/internal/config"
	"microShop/product/rpc/productclient"
	"net/url"

	"microShop/product/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Category  model.CategoryModel
	Product   model.ProductModel
	Recommend model.RecommendModel
	BizRedis  *redis.Redis
	Rpc       productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:    c,
		Category:  model.NewCategoryModel(sqlx.NewMysql(DataSource), c.CacheRedis),
		Product:   model.NewProductModel(sqlx.NewMysql(DataSource), c.CacheRedis),
		Recommend: model.NewRecommendModel(sqlx.NewMysql(DataSource), c.CacheRedis),
		BizRedis:  redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		Rpc:       productclient.NewProduct(zrpc.MustNewClient(c.Rpc)),
	}
}
