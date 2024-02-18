package logic

import (
    "context"
    "fmt"

    "add/internal/svc"
    "add/protoc/add"

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
    val00 := l.ctx.Value("Key00")
    fmt.Printf("val00: %v", val00)

    var s add.AddResp
    s.Sum = in.A + in.B
    fmt.Printf("%d + %d = %d\n", in.A, in.B, s.Sum)
    //return &add.AddResp{}, nil
    return &s, nil
}
