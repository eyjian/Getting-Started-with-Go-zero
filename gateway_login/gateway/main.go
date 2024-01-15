package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/middleware"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	var c gateway.GatewayConf
	flag.Parse()

	conf.MustLoad(*configFile, &c)
	server := gateway.MustNewServer(c)
	server.Use(middleware.LoginAndAuthMiddleware)
	defer server.Stop()

	// 实例化登录服务客户端
	middleware.NewLoginClient()

	// 设置错误处理
	httpx.SetErrorHandlerCtx(grpcErrorHandlerCtx)

	fmt.Printf("Starting gateway at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func grpcErrorHandler(err error) (int, any) {
	if st, ok := status.FromError(err); ok {
		return http.StatusOK, MyResponse{
			Code:    int(st.Code()),
			Message: st.Message(),
		}
	}

	code := 2024
	return http.StatusOK, MyResponse{
		Code:    code,
		Message: err.Error(),
	}
}

func grpcErrorHandlerCtx(ctx context.Context, err error) (int, any) {
	return grpcErrorHandler(err)
}

type MyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
