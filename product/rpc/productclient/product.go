// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"microShop/product/rpc/types/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResply   = product.CommonResply
	GetCateoryReq  = product.GetCateoryReq
	GetProductsReq = product.GetProductsReq
	ProductReq     = product.ProductReq

	Product interface {
		Products(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (*CommonResply, error)
		Product(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*CommonResply, error)
		Category(ctx context.Context, in *GetCateoryReq, opts ...grpc.CallOption) (*CommonResply, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Products(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Products(ctx, in, opts...)
}

func (m *defaultProduct) Product(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Product(ctx, in, opts...)
}

func (m *defaultProduct) Category(ctx context.Context, in *GetCateoryReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Category(ctx, in, opts...)
}
