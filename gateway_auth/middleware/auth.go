package middleware

import (
    "context"
    "errors"
    "net/http"

    //"github.com/zeromicro/go-zero/rest"
)

var ErrInvalidToken = errors.New("invalid token")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        // 调用登录服务接口生成鉴权信息
        // 这里只是一个示例，你需要替换为实际的登录服务调用
        if token != "valid-token" {
            //rest.WriteError(w, r, ErrInvalidToken)
            w.Write([]byte("invalid token"))
            return
        }

        // 将鉴权信息添加到请求上下文中
        ctx := context.WithValue(r.Context(), "token", token)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

