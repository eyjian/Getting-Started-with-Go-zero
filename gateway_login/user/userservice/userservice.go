// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"user/protoc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserReq  = user.UserReq
	UserResp = user.UserResp

	UserService interface {
		QueryUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) QueryUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.QueryUser(ctx, in, opts...)
}
