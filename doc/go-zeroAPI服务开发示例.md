### 接口定义

* **定义 API 接口文件**

接口文件 add.api 的内容如下：

```
syntax = "v1"

info (
    title:   "API 接口文件示例"
    desc:    "演示如何编写 API 接口文件"
    author:  "一见"
    date:    "2023年12月07日"
    version: "v1"
)

type AddReq {
    A int `path:"a"`
    B int `path:"b"`
}

type AddReply {
    Sum int `json:"sum"`        
}

service Adder {
    @handler add
    get /add/:a/:b(AddReq) returns(AddReply)
}

上述go-zero的api文件定义语法正确吗？
```

* **编译 API 接口文件**

在 add.api 文件所在目录下，使用 go-zero 的脚手架工具 goctl 编译 add.api 文件：

```sh
# goctl api go -api add.api -dir .
Done.
```

编译成功后的目录结构：

```
.
├── add.api
├── adder.go # 服务端 main 函数所在文件
├── etc
│   └── adder.yaml # 配置文件
└── internal
    ├── config
    │   └── config.go # 和配置对应的数据结构
    ├── handler # HTTP 部分代码
    │   ├── addhandler.go
    │   └── routes.go
    ├── logic
    │   └── addlogic.go # 需要实现的业务逻辑代码
    ├── svc
    │   └── servicecontext.go # 上下文
    └── types
        └── types.go # 对应 API 中定义的数据结构
```

在进一步之前，还需执行“go mod tidy”整理依赖。

目录 etc 下的配置文件 adder.yaml 定义的 API 网关服务的服务端口等：

```
# cat etc/adder.yaml 
Name: Adder
Host: 0.0.0.0
Port: 8888
```

如上所示，go-zero 脚手架 goctl 设置的监听端口为 8888 。

* **编译生成可执行程序文件**

```
go mod tidy
go build -o add_http_server adder.go
```

* **启动服务：**

```
# ./add_http_server 
Starting server at 0.0.0.0:8888...
```

* **测试服务是否可用：**

```
# curl -i "http://localhost:8888/add"
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-97de89193a15ff3704beeab6ab01cbc5-448ad910b934da13-00
Date: Thu, 07 Dec 2023 07:09:36 GMT
Content-Length: 4

null
```

### 接口实现

go-zero 的脚手架 goctl 生成的是一个空服务，除了返回 null，啥也没干。在函数 Add 中增加实现：

```go
# cat internal/logic/addlogic.go 
package logic

import (
        "context"

        "api/internal/svc"
        "api/internal/types"

        "github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
        logx.Logger
        ctx    context.Context
        svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
        return &AddLogic{
                Logger: logx.WithContext(ctx),
                ctx:    ctx,
                svcCtx: svcCtx,
        }
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddReply, err error) {
        // todo: add your logic here and delete this line
        s := req.A + req.B // 新增代码
        return &types.AddReply{ s }, nil // 新增代码
        return
}
```

重新编译执行：

```sh
# curl -i "http://localhost:8888/add?a=1&b=3"
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-33930e740d4d642dd8a16667af5f6a60-cb450c0bdbb9f6fc-00
Date: Fri, 08 Dec 2023 02:58:48 GMT
Content-Length: 9

{"sum":4}
```
