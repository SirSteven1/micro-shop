package logic

import (
	"context"

	"microShop/home/api/internal/svc"
	"microShop/home/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerLogic {
	return &BannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BannerLogic) Banner(req *types.BannerReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
