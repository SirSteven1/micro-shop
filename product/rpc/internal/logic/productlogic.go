package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"microShop/product/rpc/internal/svc"
	"microShop/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductReq) (*product.CommonResply, error) {
	id, _ := strconv.ParseInt(in.ProductIds, 0, 64)
	res, err := l.svcCtx.Product.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}
	jsonRes, _ := json.Marshal(res)
	return &product.CommonResply{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    string(jsonRes),
	}, nil
}
