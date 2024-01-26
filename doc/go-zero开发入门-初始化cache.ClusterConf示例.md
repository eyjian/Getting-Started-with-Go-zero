cache.ClusterConf 的定义如下：

```go
// CacheConf is an alias of ClusterConf.
type CacheConf = ClusterConf

type (
	// A ClusterConf is the config of a redis cluster that used as cache.
	ClusterConf []NodeConf

	// A NodeConf is the config of a redis node that used as cache.
	NodeConf struct {
		redis.RedisConf
		Weight int `json:",default=100"`
	}
)

type (
	// A RedisConf is a redis config.
	RedisConf struct {
		Host     string
		Type     string `json:",default=node,options=node|cluster"`
		Pass     string `json:",optional"`
		Tls      bool   `json:",optional"`
		NonBlock bool   `json:",default=true"`
		// PingTimeout is the timeout for ping redis.
		PingTimeout time.Duration `json:",default=1s"`
	}
)
```

初始化示例：

```go
cacheConf := cache.ClusterConf{
	{ // 数组
		RedisConf: redis.RedisConf{
			Host: r1.Addr,
			Type: redis.NodeType,
		},
		Weight: 100,
	},
}
```

配置文件示例：

```yaml
CacheRedis: 
    - Host: 127.0.0.2
      port: 6379
      type: node
```

配置文件对应的源代码：

```go
type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
}

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("Host: %s\n", c.CacheRedis[0].Host)
}
```
