package logic

import (
	"context"
	"errors"
	"fmt"

	"authentication/internal/svc"
	"authentication/protoc/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthenticateLogic) Authenticate(in *auth.AuthReq) (*auth.AuthResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("SessionId: %s\n", in.SessionId)
	if in.SessionId == "2023121200" {
		authResp := &auth.AuthResp {
			UserId: "10001",
		}
		fmt.Printf("UserId: %s\n", authResp.UserId)
		return authResp, nil
	} else {
		fmt.Printf("no perm: %s\n", in.SessionId)
		return nil, errors.New("No perm in auth.Authenticate")
	}
}
