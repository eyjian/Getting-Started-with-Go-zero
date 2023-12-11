go-zero开发入门-API网关开发示例

2023/12/08

开发一个 API 网关，代理 [https://blog.csdn.net/Aquester/article/details/134856271](https://blog.csdn.net/Aquester/article/details/134856271) 中的 RPC 服务。

### 网关完整源代码

```go
// file: main.go
package main

import (
    "flag"
    "fmt"

    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
    var c gateway.GatewayConf
    flag.Parse()
    
    conf.MustLoad(*configFile, &c)
    server := gateway.MustNewServer(c)
    defer server.Stop()

    fmt.Printf("Starting gateway on %s:%d...\n", c.Host, c.Port)
    server.Start()
}
```

上述代码可以使用 go-zero 的脚手架工具 goctl 自动生成，比如：

```
goctl gateway -dir gateway
```

同时会生成网关配置文件 gateway.yaml：

```yaml
# cat etc/gateway.yaml 
Name: gateway-example # gateway name
Host: localhost # gateway host
Port: 8888 # gateway port
Upstreams: # upstreams
  - Grpc: # grpc upstream
      Target: 0.0.0.0:8080 # grpc target,the direct grpc server address,for only one node
#      Endpoints: [0.0.0.0:8080,192.168.120.1:8080] # grpc endpoints, the grpc server address list, for multiple nodes
#      Etcd: # etcd config, if you want to use etcd to discover the grpc server address
#        Hosts: [127.0.0.1:2378,127.0.0.1:2379] # etcd hosts
#        Key: greet.grpc # the discovery key
    # protoset mode
    ProtoSets:
      - hello.pb
    # Mappings can also be written in proto options
#    Mappings: # routes mapping
#      - Method: get
#        Path: /ping
#        RpcPath: hello.Hello/Ping
```

### 编译网关源码生成可执行程序文件

```
# cat Makefile 
all: gateway
gateway: main.go
        go build -o gateway main.go 

clean:
        rm -f gateway
```

### 生成被代理 RPC 服务的 pb 文件

```sh
protoc --descriptor_set_out=add.pb add.proto
```

add.pb 是一个二进制文件。

### 编辑网关配置文件 gateway.yaml

```
# cat etc/gateway.yaml 
Name: go-zero-gateway # 网关名
Host: 0.0.0.0 # 网关的服务地址
Port: 8888 # 网关的服务端口
Upstreams: # 被网关代理的上游服务
  - Grpc: # gRPC 服务列表
      Etcd:
        Hosts:
        - 127.0.0.1:2379 # etcd 服务地址和端口（Endpoints）
        Key: add.rpc
    ProtoSets: # 这里为被代理 RPC 服务的 pb 文件列表
        - /root/go-zero/gateway/proto/add.pb
    Mappings: # Mappings can also be written in proto options
      - Method: get # HTTP 方法
        Path: /add # HTTP 路径
        RpcPath: add.Adder/add # 对应的 RPC 路径（第一个 add 为包名，Adder 为 service 名，后一个 add 为 service 中的方法名）
```

### 启动网关

```
./gateway
```

### 通过网关访问 RPC 服务

```
# curl -i '127.0.0.1:8888/add'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-a53f71dd9eb638c5c2af03eb633a56be-dad8426820c63b1d-00
Date: Fri, 08 Dec 2023 02:08:39 GMT
Content-Length: 9

{"sum":0}

# curl -i '127.0.0.1:8888/add?a=1&b=2'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-0df2bbd651938b704c532a01bb2f16e3-27c5267777c8cc06-00
Date: Fri, 08 Dec 2023 02:08:47 GMT
Content-Length: 9

{"sum":3}
```
