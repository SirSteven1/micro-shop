package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"microShop/user/api/internal/svc"
	"microShop/user/api/internal/types"
	"microShop/user/rpc/types/account"

	"microShop/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.LoginReq) (resp *types.CommonResply, err error) {
	userIdentity := fmt.Sprintf("%v", l.ctx.Value("userIdentity"))

	userData, userErr := l.svcCtx.Rpc.UserInfo(l.ctx, &account.UserInfoReq{UserIdentity: userIdentity})
	if userErr != nil {
		return nil, userErr
	}

	userInfoItem := types.UserInfoItem{}

	userJsonErr := json.Unmarshal([]byte(userData.Data), &userInfoItem)
	if userJsonErr != nil {
		return nil, errorx.NewDefaultError("数据转换失败")
	}
	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    &userInfoItem,
	}, nil
}
