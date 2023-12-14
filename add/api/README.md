### 编译步骤

以下均在本项目的根目录下执行，注意不是指 Getting-Started-with-Go-zero 的根目录。

* 编译 api 文件：

```shell
make api
```

生成可执行程序文件 add_http_server：

```shell
make
```

### 启动运行

```shell
./add_http_server
```

### 测试验证

```shell
curl -i "http://localhost:8888/add"
```
