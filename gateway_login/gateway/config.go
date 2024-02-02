package main

import (
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	gateway.GatewayConf
	Login zrpc.RpcClientConf
}
