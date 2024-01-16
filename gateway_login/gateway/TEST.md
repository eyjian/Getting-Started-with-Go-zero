测试用例

```sh
curl -i '127.0.0.1:9001/v1/login'
curl -i '127.0.0.1:9001/v1/login?phone=138&verification_code=8899'

curl -b "mysid=12345" -i '127.0.0.1:9001/v2/query_user' # 鉴权失败
curl -b "mysid=2023121200" -i '127.0.0.1:9001/v2/query_user?uid=98&token=123' # 返回成功
curl -b "mysid=2023121200" -i '127.0.0.1:9001/v2/query_user?uid=0&token=1223' # 返回错误

# POST 也可以，但 gateway 的 Mappings.Mathod 值需为 post
curl -i -b "mysid=2023121200" -H "Content-Type: application/json" -X POST -d '{"uid":98,"token":"123"}' '127.0.0.1:9001/v2/query_user'

curl -b "mysid=2023121200" -i '127.0.0.1:9001/v3/query_user?uid=98&token=123'
```
