package logic

import (
	"context"

	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *cart.AddCartReq) (*cart.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &cart.CommonResply{}, nil
}
