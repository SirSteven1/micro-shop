package svc

import (
	"fmt"
	"microShop/product/model"
	"microShop/product/rpc/internal/config"
	"net/url"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	Product  model.ProductModel
	Category model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:   c,
		Product:  model.NewProductModel(sqlx.NewMysql(DataSource), c.Cache),
		Category: model.NewCategoryModel(sqlx.NewMysql(DataSource), c.Cache),
	}
}
