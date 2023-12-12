package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"

	"gateway/authclient"
	"gateway/loginclient"
	"gateway/protoc/auth"
	"gateway/protoc/login"
)

// 登录鉴权
func LoginAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.URL.Path: %s\n", r.URL.Path)

		if r.URL.Path == "/login" { // 登录
			LoginMiddleware(next, w, r)
		} else { // 非登录
			AuthMiddleware(next, w, r)

			// 请求放行给下游业务服务
			//            next.ServeHTTP(w, r)
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
	vars := mux.Vars(r) // 反序列化 url 参数

	loginReq.Phone = vars["phone"]
	loginReq.VerificationCode = vars["vcode"]
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
		authResp, err := authClient.Authenticate(r.Context(), &authReq)
		if err != nil {
			// 未通过鉴权
			fmt.Println(err)
			fmt.Fprintln(w, err)
		} else {
			// 通过鉴权
			//w.Header().Set("myuid", authResp.UserId)
			newReq := r.WithContext(r.Context())
			newReq.Header.Set("myuid", authResp.UserId)
			next.ServeHTTP(w, newReq)
		}
	}
}
