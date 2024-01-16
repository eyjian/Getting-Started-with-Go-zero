package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/status"
	"net/http"
	"net/url"
	"strings"

	"gateway/authclient"
	"gateway/loginclient"
	"gateway/protoc/auth"
	"gateway/protoc/login"
)

// 登录服务客户端
var loginClient loginclient.Login

// NewLoginClient 实例化登录服务客户端
func NewLoginClient() {
	var loginConf zrpc.RpcClientConf

	conf.MustLoad("etc/login.yaml", &loginConf)
	// 注意如果这里失败会直接退出，因此在目标服务不可用时，总是退出。
	// 如果不想退出，使用 zrpc.NewClient 替代 zrpc.MustNewClient。
	// 实际上 zrpc.MustNewClient 只是包了下 zrpc.NewClient，增加了 Must 逻辑。
	client := zrpc.MustNewClient(loginConf)
	loginClient = loginclient.NewLogin(client)
}

type MyError struct {
	Code uint32 `json:"code"`
	Message string `json:"message"`
}

// LoginMiddleware 登录
func LoginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[LoginMiddleware] r.Body ==> %s\n",r.Body)
		fmt.Printf("[LoginMiddleware] r.URL.RawQuery: %s\n", r.URL.RawQuery)

		if !strings.HasPrefix(r.URL.Path, "/v1/") {
			next.ServeHTTP(w, r)
		} else {
			var loginReq login.LoginReq

			params, _ := url.ParseQuery(r.URL.RawQuery)
			loginReq.Phone = params.Get("phone")
			loginReq.VerificationCode = params.Get("vcode")
			fmt.Printf("Phone:%s, VerificationCode:%s\n", loginReq.Phone, loginReq.VerificationCode)
			loginResp, err := loginClient.Login(r.Context(), &loginReq)
			if err != nil {
				var me MyError
				fmt.Println("login fail")
				//fmt.Fprintln(w, err)
				st, ok := status.FromError(err)
				if ok {
					fmt.Printf("%d", st.Code())
					me.Code = uint32(st.Code())
					me.Message = st.Message()
				} else {
					me.Code = 888899
					me.Message = err.Error()
				}
				jsonStr, _ := json.Marshal(&me)
				fmt.Fprintln(w, string(jsonStr))
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
	}
}

// AuthMiddleware 鉴权
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[AuthMiddleware] r.Body ==> %s\n",r.Body)
		fmt.Printf("[AuthMiddleware] r.URL.RawQuery: %s\n", r.URL.RawQuery)

		if !strings.HasPrefix(r.URL.Path, "/v2/") {
			next.ServeHTTP(w, r)
		} else {
			var me MyError

			cookie, err := r.Cookie("mysid")
			if err != nil {
				// cookies 中无会话 ID
				fmt.Println(w, "no access in gateway.AuthMiddleware")
				//fmt.Fprintln(w, "no access in gateway.AuthMiddleware")
				me.Code = 666699
				me.Message = "no access in gateway.AuthMiddleware"
				jsonStr, _ := json.Marshal(&me)
				fmt.Fprintln(w, string(jsonStr))
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
					//fmt.Fprintln(w, err)
					st, ok := status.FromError(err)
					if ok {
						fmt.Printf("%d", st.Code())
						me.Code = uint32(st.Code())
						me.Message = st.Message()
					} else {
						me.Code = 999988
						me.Message = err.Error()
					}
					jsonStr, _ := json.Marshal(&me)
					fmt.Fprintln(w, string(jsonStr))
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
	}
}
