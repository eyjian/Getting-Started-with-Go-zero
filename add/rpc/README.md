### 编译步骤

以下均在本项目的根目录下执行，注意不是指 Getting-Started-with-Go-zero 的根目录。

* 编译 proto 文件：

```shell
make rpc
```

生成可执行程序文件 add_server 和 add_client：

```shell
make
```

### 启动运行

* 启动服务端：

```shell
./add_server
```

* 运行客户端：

```shell
./add_client
```

### 测试验证

处理使用客户端，也可用 grpcurl 测试验证：

```shell
grpcurl -plaintext -d '{"a": 1, "b": 2}' 127.0.0.1:8080 add.Adder/add
```
