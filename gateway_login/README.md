
* **gateway**

网关

* **login**

登录服务

* **authentication**

鉴权服务

* **user**

用户服务

* **启动顺序**

```
login_service
auth_service
user_service
gateway_login
```

请求输出示例（Method 配置为 POST）：

```
# curl -i -b "mysid=2023121200" -H "Content-Type: application/json" -X POST -d '{"uid":98,"token":"123"}' '127.0.0.1:9001/v2/query_user'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-31e006d2adf790144f0144fdc761fd6d-98d39d8c929ac627-00
Date: Fri, 19 Jan 2024 06:13:24 GMT
Content-Length: 108

{"code":0,"data":{"age":23,"gender":"FEMALE","token":"123","uid":98,"uname":"zhangsan"},"message":"success"}
```

![部署结构](https://github.com/eyjian/Getting-Started-with-Go-zero/blob/main/gateway_login/deploy.png)
