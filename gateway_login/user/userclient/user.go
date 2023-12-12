// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"user/protoc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserReq  = user.UserReq
	UserResp = user.UserResp

	User interface {
		QueryUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) QueryUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryUser(ctx, in, opts...)
}