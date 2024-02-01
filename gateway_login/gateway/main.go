package main

import (
	"bytes"
	"context"
	"encoding/json"
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
	var c Config
	//var loginConf zrpc.RpcClientConf
	flag.Parse()

	conf.MustLoad(*configFile, &c)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Mode: %s\n", c.Mode)
	fmt.Printf("Host: %s\n", c.Host)
	fmt.Printf("Port: %d\n\n", c.Port)

	//conf.MustLoad("etc/login.yaml", &loginConf)
	//fmt.Printf("Etcd.Hosts: %s\n", loginConf.Etcd.Hosts)
	//fmt.Printf("Etcd.Key: %s\n", loginConf.Etcd.Key)
	fmt.Printf("Login.Etcd.Hosts: %s\n", c.Login.Etcd.Hosts)
	fmt.Printf("Login.Etcd.Key: %s\n", c.Login.Etcd.Key)
	fmt.Printf("Login.Target: %s\n", c.Login.Target)
	fmt.Printf("Login.Endpoints: %v\n\n", c.Login.Endpoints)

	server := gateway.MustNewServer(c.GatewayConf)
	server.Use(middleware.LoginMiddleware)
	server.Use(middleware.AuthMiddleware)
	server.Use(wrapResponse)
	defer server.Stop()

	// 实例化登录服务客户端
	//middleware.NewLoginClient(loginConf)
	middleware.NewLoginClient(c.Login)

	// 设置成功处理
	//httpx.SetOkHandler(grpcOkHandler)

	// 设置错误处理
	httpx.SetErrorHandlerCtx(grpcErrorHandlerCtx)

	fmt.Printf("Starting gateway at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(p []byte) (int, error) {
	return rw.body.Write(p)
}

func (rw *responseWriter) Body() []byte {
	return rw.body.Bytes()
}

// 对响应加上“"code":0,"data":{}”，
// 对于已经包含了“code”的不做任何处理（原因是 grpcErrorHandler 才能处理好）
func wrapResponse(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 记录原始响应 writer
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// 执行下一个中间件或处理函数
		next.ServeHTTP(rw, r)

		// 检查响应状态码
		if rw.statusCode != http.StatusOK {
			return
		}

		// 获取原始响应数据
		var resp map[string]interface{}
		err := json.Unmarshal(rw.Body(), &resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 检查响应是否已经包含 code
		if _, ok := resp["code"]; ok {
			// 如果响应已经包含 code，则直接写回原始响应正文
			w.Header().Set("Content-Type", "application/json")
			w.Write(rw.Body())
			return
		}

		// 包装响应数据
		wrappedResp := map[string]interface{}{
			"code":    0,
			"message": "success",
			"data":    resp,
		}

		// 将包装后的响应数据写回 response  body
		//w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(wrappedResp)
		httpx.OkJson(w, wrappedResp) // 这里的实现不要有调用 httpx.SetOkHandler
	})
}

func grpcOkHandler(ctx context.Context, a any) any {
	fmt.Printf("OKHandler => %s\n", a)
	return MyResponse{
		Code:    0,
		Message: "success",
		Data:    a,
	}
}

func grpcErrorHandler(err error) (int, any) {
	fmt.Printf("ErrorHandler => %s\n", err)
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
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
