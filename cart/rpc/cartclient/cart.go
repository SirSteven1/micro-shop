// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package cartclient

import (
	"context"

	"microShop/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCartReq   = cart.AddCartReq
	CommonResply = cart.CommonResply
	DelCartReq   = cart.DelCartReq
	EditCartReq  = cart.EditCartReq

	Cart interface {
		AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*CommonResply, error)
		EditCart(ctx context.Context, in *EditCartReq, opts ...grpc.CallOption) (*CommonResply, error)
		DelCart(ctx context.Context, in *DelCartReq, opts ...grpc.CallOption) (*CommonResply, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

func (m *defaultCart) AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.AddCart(ctx, in, opts...)
}

func (m *defaultCart) EditCart(ctx context.Context, in *EditCartReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.EditCart(ctx, in, opts...)
}

func (m *defaultCart) DelCart(ctx context.Context, in *DelCartReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.DelCart(ctx, in, opts...)
}