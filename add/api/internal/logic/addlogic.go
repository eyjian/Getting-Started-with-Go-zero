package logic

import (
    "api/adder"
    "context"
    "fmt"

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
    ctx := l.ctx
    ctx = context.WithValue(ctx, "Key00", "value00")
    ctx = context.WithValue(ctx, "Key01", "value01")
    ctx = context.WithValue(ctx, "Grpc-Metadata-Key02", "value02")
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
