配置定义：

```go
# cat internal/config/config.go 
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
}
```

对应的配置文件：

```yaml
# cat etc/abc-api.yaml 
Name: abc-api
Host: 0.0.0.0
Port: 8888

CacheRedis: 
    - Host: 127.0.0.2
      port: 6379
      type: node
```

加载配置：

```go
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("Host: %s\n", c.CacheRedis[0].Host)
}
```

执行效果：

```
# ./abc
Host: 127.0.0.2
Type: node
```

[配置源码](https://github.com/zeromicro/go-zero/blob/master/core/stores/redis/conf.go)

```go
// CacheConf is an alias of ClusterConf.
type CacheConf = ClusterConf

// A ClusterConf is the config of a redis cluster that used as cache.
type ClusterConf []NodeConf

// A NodeConf is the config of a redis node that used as cache.
type NodeConf struct {
	redis.RedisConf
	Weight int `json:",default=100"`
}

// A RedisConf is a redis config.
type RedisConf struct {
	Host     string
	Type     string `json:",default=node,options=node|cluster"`
	Pass     string `json:",optional"`
	Tls      bool   `json:",optional"`
	NonBlock bool   `json:",default=true"`
	// PingTimeout is the timeout for ping redis.
	PingTimeout time.Duration `json:",default=1s"`
}
```
