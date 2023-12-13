package middleware

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"net/url"
	"strings"

	"gateway/authclient"
	"gateway/loginclient"
	"gateway/protoc/auth"
	"gateway/protoc/login"
)

// 登录和鉴权
func LoginAndAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.URL.Path: %s\n", r.URL.Path)

		if strings.HasPrefix(r.URL.Path, "/v1/") {
			LoginMiddleware(next, w, r) // 登录请求
		} else if strings.HasPrefix(r.URL.Path, "/v2/") {
			AuthMiddleware(next, w, r) // 需鉴权的请求
		} else {
			OtherMiddleware(next, w, r) // 其它请求
		}
	}
}

// 登录
func LoginMiddleware(next http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	var loginReq login.LoginReq
	var loginConf zrpc.RpcClientConf

	conf.MustLoad("etc/login.yaml", &loginConf)
	client := zrpc.MustNewClient(loginConf)
	loginClient := loginclient.NewLogin(client)
	params, _ := url.ParseQuery(r.URL.RawQuery)

	loginReq.Phone = params.Get("phone")
	loginReq.VerificationCode = params.Get("vcode")
	fmt.Printf("Phone:%s, VerificationCode:%s\n", loginReq.Phone, loginReq.VerificationCode)
	loginResp, err := loginClient.Login(r.Context(), &loginReq)
	if err != nil {
		fmt.Println("login fail")
		fmt.Println(err)
	} else {
		cookie := &http.Cookie{ // 构造 cookie
			Name:  "mysid",
			Value: loginResp.SessionId,
			Path:  "/",
		}
		http.SetCookie(w, cookie) // 写 cookie
		fmt.Fprintln(w, "Cookie has been set")
	}
}

// 鉴权
func AuthMiddleware(next http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("mysid")
	if err != nil {
		// cookies 中无会话 ID
		fmt.Fprintln(w, "no acesss in gateway.AuthMiddleware")
	} else {
		// cookies 中有会话 ID
		var authReq auth.AuthReq
		var authConf zrpc.RpcClientConf

		fmt.Printf("cookie[mysid]: %s\n", cookie.Value)

		conf.MustLoad("etc/auth.yaml", &authConf)
		client := zrpc.MustNewClient(authConf)
		authClient := authclient.NewAuth(client)

		authReq.SessionId = cookie.Value
		authResp, err := authClient.Authenticate(r.Context(), &authReq) // 调用鉴权服务
		if err != nil {
			// 未通过鉴权
			fmt.Println(err)
			fmt.Fprintln(w, err)
		} else {
			// 通过鉴权
			fmt.Printf("[authResp.UserId] ==> %s\n", authResp.UserId)

			newReq := r.WithContext(r.Context())
			newReq.Header.Set("Grpc-Metadata-myuid", authResp.UserId)

			// 往下转发
			next.ServeHTTP(w, newReq)
		}
	}
}

// 其它请求
func OtherMiddleware(next http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	next.ServeHTTP(w, r)
}
