// go build -o add_client add_client.go 
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

