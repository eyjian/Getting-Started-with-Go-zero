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
    server.Use(middleware.LoginAuthMiddleware)
    defer server.Stop()

    fmt.Printf("Starting gateway at %s:%d...\n", c.Host, c.Port)
    server.Start()
}

