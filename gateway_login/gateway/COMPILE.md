第一次编译时的执行顺序：

```sh
# make rpc # 生成网关框架代码
# make pb
# make # 生成可执行程序文件
```

如果修改了 proto/login.proto 或者 proto/user.proto，需要重新执行“make pb”。
