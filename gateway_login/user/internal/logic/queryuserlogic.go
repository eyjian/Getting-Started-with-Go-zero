package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
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
	var userResp user.UserResp
	var uid string

	//uid := l.ctx.Value("myuid")
	vals := metadata.ValueFromIncomingContext(l.ctx, "myuid")
	if len(vals) == 0 {
		fmt.Printf("Can not get myuid from metadata\n")
		return nil, errors.Errorf("Can not get myuid from metadata")
	} else {
		uid = vals[0]
		fmt.Printf("vals[0] => %s\n", uid)

		userResp.Uname = "zhangsan"
		userResp.Age = age
		userResp.Gender = user.Gender_FEMALE
		fmt.Println(userResp)
		age = age + 1
		return &userResp, nil
	}
}
