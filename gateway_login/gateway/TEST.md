测试用例

```sh
curl -b "mysid=12345" -i '127.0.0.1:9001/v2/query_user' # 鉴权失败
curl -b "mysid=2023121200" -i '127.0.0.1:9001/v2/query_user?uid=98&token=123' # 返回成功
curl -b "mysid=2023121200" -i '127.0.0.1:9001/v2/query_user?uid=0&token=1223' # 返回错误
```
