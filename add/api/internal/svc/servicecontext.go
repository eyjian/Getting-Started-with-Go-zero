package svc

import (
    "api/internal/config"
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config config.Config
    Client zrpc.Client
}

func NewServiceContext(c config.Config, client zrpc.Client) *ServiceContext {
    return &ServiceContext{
        Config: c,
        Client: client,
    }
}
