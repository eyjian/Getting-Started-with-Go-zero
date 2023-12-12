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
	uid := l.ctx.Value("myuid")
	if uid != nil {
		fmt.Println("User ID:", uid)
	} else {
		fmt.Println("User ID not found in context")
	}
/*
	vals := metadata.ValueFromIncomingContext(l.ctx, "myuid")
	if len(vals) > 0 {
		fmt.Printf("vals[0] => %s\n", vals[0])
   		uid := vals[0]
   		fmt.Println("uid:", uid)
	}
	fmt.Printf("[ctx] ==> %+v\n", l.ctx)
*/

	var userResp user.UserResp
	userResp.Uname = "zhangsan"
	userResp.Age = age
	userResp.Gender = user.Gender_FEMALE
	fmt.Println(userResp)
	age = age + 1
	return &userResp, nil
}
