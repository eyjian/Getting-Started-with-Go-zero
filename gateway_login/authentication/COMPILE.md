第一次编译时的执行顺序：

```sh
# make rpc # 生成 rpc 服务框架代码
# make tidy # 整理依赖
# make # 生成可执行程序文件
```

如果修改了 authentication.proto，需要重新执行“make rpc”和“make tidy”。
