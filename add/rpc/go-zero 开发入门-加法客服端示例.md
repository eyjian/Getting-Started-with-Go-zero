go-zero 开发入门-加法客服端示例

2023/12/07

### 定义 RPC 接口文件

接口文件 add.proto 的内容如下：

```proto
syntax = "proto3";
package add;

// 当 protoc-gen-go 版本大于 1.4.0 时需加上 go_package，否则编译报错“unable to determine Go import path for”
option go_package = "./add";

message AddReq {
    int32 a = 1;
    int32 b = 2;
}

message AddResp {
    int32 sum = 1;
}

service Adder {
    rpc add(AddReq) returns(AddResp);
}
```

接口文件 add.proto 可放在项目的根目录下。

### 编译 RPC 接口文件

在 add.proto 文件所在目录下，使用 go-zero 的脚手架工具 goctl 编译 add.proto 文件：

```sh
# goctl rpc protoc add.proto --go_out=./protoc --go-grpc_out=./protoc --zrpc_out=.
Done.
```

编译成功后的目录结构：

```
.
├── adder
│   └── adder.go # 客户端直接可使用的 SDK 代码
├── add.go # 服务端 main 函数所在文件
├── add.proto # RPC 接口文件
├── etc
│   └── add.yaml # 配置文件
├── go.mod
├── internal
│   ├── config # 存放配置对应的数据结构
│   │   └── config.go
│   ├── logic # 业务逻辑代码放在这个目录下
│   │   └── addlogic.go
│   ├── server # RPC 服务端代码
│   │   └── adderserver.go
│   └── svc # 上下文代码
│       └── servicecontext.go
└── protoc
    └── add
        ├── add_grpc.pb.go # protoc 生成的 gRPC 代码
        └── add.pb.go # protoc 生成的 proto 代码
```

目录下原只有文件 add.proto，编译成功后产生了多个新的文件和目录。在进一步之前，还需执行“go mod tidy”整理依赖。

### RPC 服务端开发

* **编辑文件 addlogic.go：**

```go
package logic

import (
    "context"
    "fmt" // 新增的

    "add/internal/svc" // add 为 go.mod 中的 module 名，internal 为 go.mod 所在目录下的子目录，svc 为 internal 的子目录
    "add/protoc/add"

    "github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
    ctx context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
    return &AddLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
    // todo: add your logic here and delete this line
    var s add.AddResp // 新增的
    s.Sum = in.A + in.B // 新增的
    fmt.Printf("%d + %d = %d\n", in.A, in.B, s.Sum) // 新增的
    //return &add.AddResp{}, nil // 删除的
    return &s, nil // 新增的
}
```

* **编译生成可执行程序文件：**

```sh
go build -o add_server add.go
```

* **启动服务端：**

```sh
# ./add_server 
Starting rpc server at 0.0.0.0:8080...
```

* **使用 grpcurl 测试：**

```sh
grpcurl -plaintext -d '{"a": 1, "b": 2}' 127.0.0.1:8080 add.Adder/add
```

使用 grpcurl 的前提是开启 reflection 反射，否则执行报如下错误：

```
Failed to list services: server does not support the reflection API
```

对于 goctl 生成的，只需要在 etc 下的 yaml 配置文件增加：

```
Mode: dev
```

或者：

```
Mode: test
```

开启 reflection 的代码在根目录的 add.go 文件中：

```go
if c.Mode == service.DevMode || c.Mode == service.TestMode {
    reflection.Register(grpcServer)
}
```

### RPC 客户端开发

* **编辑客户端文件 add_client.go：**

客户端通过调用 adder/adder.go 中的函数 Add 来访问服务端，客户端代码文件 add_client.go 内容如下：

```go
package main

import (
    "context"
    "fmt"

    "add/adder"
    "add/protoc/add"

    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/zrpc"
)

func main() {
    var clientConf zrpc.RpcClientConf
    conf.MustLoad("etc/client.yaml", &clientConf)

    client := zrpc.MustNewClient(clientConf)
    adder := adder.NewAdder(client)

    addReq := &add.AddReq{ A:1, B:2 }
    addResp, err := adder.Add(context.Background(), addReq)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(addReq)
    fmt.Println(addResp)
    fmt.Println("sum=", addResp.Sum)
}
```

文件 add_client.go 也放在根目录下，和 add.go 同目录。

* **编译生成可执行程序文件：**

```sh
go build -o add_client add_client.go
```

* **执行客户端：**

```
# ./add_client
{"@timestamp":"2023-12-07T11:38:39.231+08:00","caller":"p2c/p2c.go:181","content":"p2c - conn: 127.0.0.1:8080, load: 1029, reqs: 1","level":"stat"}
sum:3
```

### 附

* **goctl 的安装参见：**

(https://blog.csdn.net/Aquester/article/details/134843086)[https://blog.csdn.net/Aquester/article/details/134843086]

* **etcd 的安装参见：**

(https://blog.csdn.net/Aquester/article/details/134843461)[https://blog.csdn.net/Aquester/article/details/134843461]

* **grpcurl 下载：**

(https://github.com/fullstorydev/grpcurl/releases)[https://github.com/fullstorydev/grpcurl/releases]

下载 x86_64 版本的 Linux 二进制包：

(https://github.com/fullstorydev/grpcurl/releases/download/v1.8.9/grpcurl_1.8.9_linux_x86_64.tar.gz)[https://github.com/fullstorydev/grpcurl/releases/download/v1.8.9/grpcurl_1.8.9_linux_x86_64.tar.gz]

grpcurl 是一个命令行工具，允许与 gRPC 服务器交互，基本上是对 gRPC 服务器的 curl 。grpcurl 的参数“-plaintext”表示不使用 TLS/SSL 进行加密通信，参数“-d”用于指定请求消息的 JSON 格式。

grpcurl 的 list 和 describe 可列出 gRPC 服务端反射的 protobuf：

```sh
grpcurl --plaintext 127.0.0.1:8080 list
```

grpcurl 服务调用格式：

```sh
grpcurl -plaintext grpc.server.com:80 my.custom.server.Service/Method
```

如果为 TLS：

```
# grpcurl grpc.server.com:443 my.custom.server.Service/Method
```

带请求参数调用格式：

```
# grpcurl -d '{"id": 1234, "tags": ["foo","bar"]}' grpc.server.com:443 my.custom.server.Service/Method
```

* **grpcui 安装：**

```sh
# go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

gRPCui 是 gRPC 的交互式 Web UI，基于 grpcurl，并提供一个 GUI 来发现和测试 gRPC 服务，类似于 Postman 或 Swagger UI 等 HTTP 工具，但是用于 gRPC API 而不是 REST。

* **升级 goctl：**

升级：

```sh
goctl env check -i -f
```

升级检查：

```
goctl env
```
