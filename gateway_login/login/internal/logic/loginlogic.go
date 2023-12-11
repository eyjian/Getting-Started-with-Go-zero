package logic

import (
	"context"
	"fmt"

	"login/internal/svc"
	"login/protoc/login"

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

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("Phone: %s, VerificationCode:%s\n",in.Phone,in.VerificationCode)
	var loginResp login.LoginResp
	loginResp.SessionId = "2023121200"
	return &loginResp, nil
}
