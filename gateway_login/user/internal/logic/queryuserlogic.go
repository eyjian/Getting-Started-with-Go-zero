package logic

import (
	"context"
	"fmt"

	"user/internal/svc"
	"user/protoc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var age int32 = 18

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
	return &QueryUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserLogic) QueryUser(in *user.UserReq) (*user.UserResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("Uid: %s\n", in.Uid)
	var userResp user.UserResp
	userResp.Uname = "zhangsan"
	userResp.Age = age
	userResp.Gender = user.Gender_FEMALE
	fmt.Println(userResp)
	age = age + 1
	return &userResp, nil
}
