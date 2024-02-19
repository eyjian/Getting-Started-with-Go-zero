客户端：

```go
md := metadata.New(map[string]string{"signature": "0123456789", "timestamp": "2021-07-01 00:00:00"})
ctx := metadata.NewOutgoingContext(ctx, md)
addResp, err := adderClient.Add(ctx, addReq) // rpc 调用
```

初始化 md 也可如下方式：

```go
md := metadata.Pairs(
	"signature", "0123456789",
	"timestamp", "2021-07-01 00:00:00",
）
ctx := metadata.NewOutgoingContext(ctx, md)
addResp, err := adderClient.Add(ctx, addReq) // rpc 调用
```

追加新的如下：

```go
ctx = metadata.AppendToOutgoingContext(ctx, "noncestr", "abc")
```

也可使用 md 的 Set 和 Append 方法追加：

```go
md.Set("extra", "extra-value")
md.Append("append", "append-value")
```

服务端:

```go
md, _ := metadata.FromIncomingContext(ctx)
或直接：
vals := metadata.ValueFromIncomingContext(ctx, "signature")
```

注意 key 都会被转为小写，即使客户端为大写：

```go
// Keys beginning with "grpc-" are reserved for grpc-internal use only and may
// result in errors if set in metadata.
func New(m map[string]string) MD {
	md := make(MD, len(m))
	for k, val := range m {
		key := strings.ToLower(k)
		md[key] = append(md[key], val)
	}
	return md
}

func Pairs(kv ...string) MD {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metadata: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}
	md := make(MD, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		key := strings.ToLower(kv[i])
		md[key] = append(md[key], kv[i+1])
	}
	return md
}

func (md MD) Set(k string, vals ...string) {
	if len(vals) == 0 {
		return
	}
	k = strings.ToLower(k)
	md[k] = vals
}
```

而且 key 只能由 数字、字母和三个特殊字符“-_.”组成，大写字母会自动被转为小写字母。
