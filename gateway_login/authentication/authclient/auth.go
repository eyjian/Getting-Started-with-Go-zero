// Code generated by goctl. DO NOT EDIT.
// Source: authentication.proto

package authclient

import (
	"context"

	"authentication/protoc/auth"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthReq  = auth.AuthReq
	AuthResp = auth.AuthResp

	Auth interface {
		Authenticate(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

func (m *defaultAuth) Authenticate(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	client := auth.NewAuthClient(m.cli.Conn())
	return client.Authenticate(ctx, in, opts...)
}
