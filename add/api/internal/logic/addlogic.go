package logic

import (
    "api/adder"
    "context"
    "fmt"
    "google.golang.org/grpc/metadata"

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
    //s := req.A + req.B
    //return &types.AddReply{Sum: s}, nil
    adderClient := adder.NewAdder(l.svcCtx.Client)
    addReq := &adder.AddReq{
        A: int32(req.A),
        B: int32(req.B),
    }

    // 两种创建 md 的方法
    //md := metadata.New(map[string]string{"signature": "0123456789", "timestamp": "2021-07-01 00:00:00"})
    md := metadata.Pairs(
        "signature", "0123456789",
        "timestamp", "2021-07-01 00:00:00",
    )
    md.Set("extra", "extra-value")
    md.Append("append", "append-value")
    ctx := metadata.NewOutgoingContext(l.ctx, md)
    ctx = metadata.AppendToOutgoingContext(ctx, "noncestr", "abc")

    addResp, err := adderClient.Add(ctx, addReq)
    if err != nil {
        fmt.Printf("Call adder.Add error: %s", err.Error())
    } else {
        fmt.Printf("Call adder.Add success: %d", addResp.Sum)
    }
    return &types.AddReply{
        Sum: int(addResp.Sum),
    }, nil
}
