go-zero开发入门-API网关鉴权开发示例

2023/12/10

本文是[go-zero开发入门-API网关开发示例](https://blog.csdn.net/Aquester/article/details/134871400)一文的延伸，继续之前请先阅读此文。

在项目根目录下创建子目录 middleware，在此目录下创建文件 auth.go，内容如下：

```go
// 鉴权中间件
package middleware

import (
    "context"
    "errors"
    "net/http"
)

var ErrInvalidToken = errors.New("invalid token")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        // 调用登录服务接口生成鉴权信息
        // 这里只是一个示例，实际需替换为实际的登录服务调用
        if token != "valid-token" {
            w.Write([]byte("invalid token"))
            return
        }

        // 将鉴权信息添加到请求上下文中
        ctx := context.WithValue(r.Context(), "token", token)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}
```

在网关 main.go 文件中加入鉴权中间件：

```go
package main

import (
    "flag"
    "fmt"
    "gateway/middleware"

    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
    var c gateway.GatewayConf
    flag.Parse()
    
    conf.MustLoad(*configFile, &c)
    server := gateway.MustNewServer(c)
    server.Use(middleware.AuthMiddleware) // 使用 server 的 Use() 方法添加全局中间件
    defer server.Stop()

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
```

在网关的配置文件 etc/gateway.yaml 中加入鉴权：

```yaml
Name: go-zero-gateway
Host: 0.0.0.0
Port: 9999
Upstreams:
  - Grpc:
      Etcd:
        Hosts:
        - 127.0.0.1:2379
        Key: add.rpc
    ProtoSets:
        - /root/go-zero/gateway_auth/proto/add.pb
    Mappings: # Mappings can also be written in proto options
      - Method: get
        Path: /add
        RpcPath: add.Adder/add
    Headers: 
        - Authorization
```

编译生成网关可执行程序文件：

```sh
# cat Makefile 
all: gateway
gateway: main.go
        go build -o gateway main.go 

clean:
        rm -f gateway

rpc:
        goctl gateway -dir gateway
```

启动网关：

```sh
./gateway
```

请求网关：

```sh
# curl -i '127.0.0.1:9999/add?a=1&b=2'
HTTP/1.1 200 OK
Traceparent: 00-1cd6f9f8c902193d8dd7da646f775d0d-4959382686bbb075-00
Date: Sun, 10 Dec 2023 12:04:40 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8

invalid token
```
