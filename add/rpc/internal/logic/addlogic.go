package logic

import (
    "add/internal/svc"
    "add/protoc/add"
    "context"
    "fmt"
    "google.golang.org/grpc/metadata"

    "github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
    ctx    context.Context
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
    md, _ := metadata.FromIncomingContext(l.ctx)
    fmt.Println(l.ctx)
    fmt.Println(md)

    signature := metadata.ValueFromIncomingContext(l.ctx, "signature")
    fmt.Printf("signature: %s\n", signature)

    myValue, ok := l.ctx.Value("MyKey").(string)
    if ok {
        fmt.Printf("MyValue: %s\n", myValue)
    } else {
        fmt.Printf("MyValue: %s\n", "error")
    }

    connectionData := l.ctx.Value("transport.connectionKey")
    fmt.Printf("connectionData: %v\n", connectionData)
    peerData := l.ctx.Value("peer.peerKey")
    fmt.Printf("peerData: %v\n", peerData)

    var s add.AddResp
    s.Sum = in.A + in.B
    fmt.Printf("%d + %d = %d\n", in.A, in.B, s.Sum)
    //return &add.AddResp{}, nil
    return &s, nil
}
