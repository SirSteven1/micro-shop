package logic

import (
	"context"

	"microShop/home/rpc/internal/svc"
	"microShop/home/rpc/types/home"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerLogic {
	return &BannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BannerLogic) Banner(in *home.BannerReq) (*home.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &home.CommonResply{}, nil
}
