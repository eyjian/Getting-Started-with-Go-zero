package logic

import (
    "context"
    "fmt"
    "github.com/pkg/errors"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "user/internal/svc"
    "user/protoc/user"

    "github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

var age int32 = 18

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
    return &QueryUserLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *QueryUserLogic) QueryUser(in *user.UserReq) (*user.UserResp, error) {
    // todo: add your logic here and delete this line
    fmt.Printf("UserReq => %s\n", in)
    var userResp user.UserResp
    var uid string

    // 输出所有的 http 头
    md, ok := metadata.FromIncomingContext(l.ctx)
    if ok {
        for key, values := range md {
            for _, value := range values {
                fmt.Printf("key:%s => %s\n", key, value)
            }
        }
    }

    // 注意是 gateway-myuid，而非 myuid，
    // 可参考 go-zero 源码（go-zero/gateway/internal/headerprocessor.go）中的实现，
    // 相关的前缀为“Grpc-Metadata-”，这是在 HTTP 的 headers 中的前缀。
    vals := metadata.ValueFromIncomingContext(l.ctx, "gateway-myuid")
    if len(vals) == 0 {
        fmt.Printf("Can not get myuid from metadata\n")
        return nil, errors.Errorf("Can not get myuid from metadata")
    } else {
        fmt.Printf("vals => %s\n", vals)
        uid = vals[0]
        fmt.Printf("vals[0] => %s\n", uid)

        if in.Uid == 0 {
            st := status.New(202621, "invalid uid")
            return nil, st.Err()
        } else {
            userResp.Uid = in.Uid
            userResp.Token = in.Token
            userResp.Uname = "zhangsan"
            userResp.Age = age
            userResp.Gender = user.Gender_FEMALE
            fmt.Println(userResp)
            age = age + 1
            return &userResp, nil
        }
    }
}
