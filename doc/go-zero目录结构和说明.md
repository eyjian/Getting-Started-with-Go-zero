```
.
├── code-of-conduct.md      行为准则
├── CONTRIBUTING.md         贡献指南
├── core                    框架的核心组件
│   ├── bloom               布隆过滤器，用于检测一个元素是否在一个集合中
│   ├── breaker             熔断器，用于防止过多的请求导致系统崩溃
│   ├── cmdline             命令行解析，提供了一个简单易用的命令行参数解析器
│   ├── codec               编解码器，提供了多种编解码方式，如 JSON、Protobuf 等
│   ├── collection          集合操作，提供了一些常用的集合操作方法
│   ├── color               颜色处理，提供了一些颜色处理方法
│   ├── conf                配置文件解析，提供了一个简单易用的配置文件解析器
│   ├── contextx            上下文扩展，提供了一些扩展标准库 context 的方法
│   ├── discov              服务发现，提供了一些服务发现的接口和实现
│   ├── errorx              错误处理，提供了一些错误处理的方法和工具
│   ├── executors           执行器，提供了一些执行任务的方法和工具
│   ├── filex               文件操作，提供了一些文件操作的方法和工具
│   ├── fs                  文件系统，提供了一些文件系统相关的方法和工具
│   ├── fx                  依赖注入，提供了一个简单易用的依赖注入框架
│   ├── hash                哈希算法，提供了一些哈希算法的实现
│   ├── iox                 I/O 操作，提供了一些 I/O 操作的方法和工具
│   ├── jsonx               JSON 操作，提供了一些 JSON 操作的方法和工具
│   ├── lang                语言扩展，提供了一些语言相关的方法和工具
│   ├── limit               限流器，提供了一些限流算法的实现
│   ├── load                负载均衡，提供了一些负载均衡算法的实现
│   ├── logc                日志钩子，提供了一些日志钩子的实现
│   ├── logx                日志扩展，提供了一些扩展标准库 log 的方法和工具
│   ├── mapping             映射操作，提供了一些映射操作的方法和工具
│   ├── mathx               数学扩展，提供了一些数学相关的方法和工具
│   ├── metric              度量指标，提供了一些度量指标的实现
│   ├── mr                  MapReduce，提供了一些 MapReduce 相关的方法和工具
│   ├── naming              命名规则，提供了一些命名规则的实现
│   ├── netx                网络操作，提供了一些网络操作的方法和工具
│   ├── proc                进程操作，提供了一些进程操作的方法和工具
│   ├── prof                性能分析，提供了一些性能分析的方法和工具
│   ├── prometheus          Prometheus 监控，提供了一些与 Prometheus 集成的方法和工具
│   ├── queue               队列操作，提供了一些队列操作的方法和工具
│   ├── rescue              异常恢复，提供了一些异常恢复的方法和工具
│   ├── search              搜索操作，提供了一些搜索操作的方法和工具
│   ├── service             服务封装，提供了一些服务封装的方法和工具
│   ├── stat                统计操作，提供了一些统计操作的方法和工具
│   ├── stores              存储操作，提供了一些存储操作的方法和工具
│   ├── stringx             字符串操作，提供了一些字符串操作的方法和工具
│   ├── syncx               同步操作，提供了一些同步操作的方法和工具
│   ├── sysx                系统操作，提供了一些系统操作的方法和工具
│   ├── threading           线程操作，提供了一些线程操作的方法和工具
│   ├── timex               时间操作，提供了一些时间操作的方法和工具
│   ├── trace               链路追踪，提供了一些链路追踪的实现
│   ├── utils               实用工具，提供了一些实用的工具方法
│   └── validation          验证操作，提供了一些验证操作的方法和工具
├── gateway                 API 网关实现，提供了一个高性能、可扩展的 API 网关
│   ├── config.go           网关配置文件实现
│   ├── internal
│   ├── readme.md
│   ├── server.go           网关的实现
│   └── server_test.go
├── go.mod                  记录了项目的 Go 模块依赖
├── go.sum                  记录了项目的 Go 模块校验和
├── internal                包含了 go-zero 的内部实现，主要包括一些测试和工具
│   ├── dbtest
│   ├── devserver
│   ├── encoding
│   ├── health
│   ├── mock
│   └── trace
├── LICENSE
├── readme-cn.md             go-zero 的中文文档
├── readme.md                go-zero 的英文文档
├── rest                     RESTful API 实现，提供了一个简单易用的方式来构建 RESTful API
│   ├── chain
│   ├── config.go            rest 服务的配置实现
│   ├── engine.go            rest 服务引擎
│   ├── engine_test.go
│   ├── handler
│   ├── httpc
│   ├── httpx
│   ├── internal
│   ├── pathvar
│   ├── router
│   ├── server.go            rest 服务实现
│   ├── server_test.go
│   ├── token
│   └── types.go
├── tools                    目前仅包含脚手架工具 goctl 的实现
│   └── goctl
└── zrpc                     gRPC 服务实现，提供了一个简单易用的方式来构建 gRPC 服务
    ├── client.go            gRPC 服务的客户端实现
    ├── client_test.go
    ├── config.go            gRPC 服务的配置实现，含服务端配置（RpcServerConf）和客户端配置（RpcClientConf）的实现
    ├── config_test.go
    ├── internal             gRPC 服务的内部实现
    ├── proxy.go             gRPC 代理服务的实现
    ├── proxy_test.go
    ├── resolver
    ├── server.go            gRPC 服务的服务端实现
    └── server_test.go
```

上述很多命名以“x”结尾，比如：contextx、errorx、filesx、logx 和 httpx 等，这里的“x”均为扩展的意思，为英文单词“eXtension”的缩写。
