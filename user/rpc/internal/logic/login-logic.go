package logic

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	comm "microShop/common"
	"microShop/user/model"
	"microShop/user/rpc/internal/svc"
	"microShop/user/rpc/types/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *account.LoginReq) (*account.CommonResply, error) {
	//获取用户信息
	userData, _ := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if userData == nil {
		return nil, errors.New("用户未注册")
	}

	//密码解码
	dePassWord, dePassErr := comm.Decrypt(userData.PassWord, []byte(l.svcCtx.Config.SecretKey))
	if dePassErr != nil {
		return nil, dePassErr
	}
	if dePassWord != in.PassWord {
		return nil, errors.New("密码错误")
	}

	// 获取客户端IP
	ip, _ := comm.ExternalIp()
	user := model.User{
		Id:           userData.Id,
		UserIdentity: userData.UserIdentity,
		UserName:     userData.UserName,
		PassWord:     userData.PassWord,
		UserNick:     userData.UserNick,
		UserFace:     userData.UserFace,
		UserSex:      userData.UserSex,
		UserEmail:    userData.UserEmail,
		UserPhone:    userData.UserPhone,
		LoginAddress: ip.String(),
	}

	cntErr := l.svcCtx.UserModel.Update(l.ctx, &user)
	if cntErr != nil {
		return nil, cntErr
	}

	userJsonData, jsonErr := json.Marshal(userData)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &account.CommonResply{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data:    string(userJsonData),
	}, nil
}
