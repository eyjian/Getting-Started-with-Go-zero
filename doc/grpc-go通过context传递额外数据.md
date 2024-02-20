* **使用 ctx.Value 从 context 读取数据**

```go
// ValueFromIncomingContext returns the metadata value corresponding to the metadata
// key from the incoming metadata if it exists. Key must be lower-case.
//
// # Experimental
//
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
// later release.
func ValueFromIncomingContext(ctx context.Context, key string) []string {
	md, ok := ctx.Value(mdIncomingKey{}).(MD)
	if !ok {
		return nil
	}

	if v, ok := md[key]; ok {
		return copyOf(v)
	}
	for k, v := range md {
		// We need to manually convert all keys to lower case, because MD is a
		// map, and there's no guarantee that the MD attached to the context is
		// created using our helper functions.
		if strings.ToLower(k) == key {
			return copyOf(v)
		}
	}
	return nil
}
```

* **使用 ctx.Value 往 context 写入数据**

```go
// AppendToOutgoingContext returns a new context with the provided kv merged
// with any existing metadata in the context. Please refer to the documentation
// of Pairs for a description of kv.
func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metadata: AppendToOutgoingContext got an odd number of input pairs for metadata: %d", len(kv)))
	}
	md, _ := ctx.Value(mdOutgoingKey{}).(rawMD)
	added := make([][]string, len(md.added)+1)
	copy(added, md.added)
	kvCopy := make([]string, 0, len(kv))
	for i := 0; i < len(kv); i += 2 {
		kvCopy = append(kvCopy, strings.ToLower(kv[i]), kv[i+1])
	}
	added[len(added)-1] = kvCopy
	return context.WithValue(ctx, mdOutgoingKey{}, rawMD{md: md.md, added: added})
}
```

metadata 是 grpc 内置的，用来往 RPC 服务传递 http 头数据，分 in 和 out 两种，对应的 key 都为一个空 struct，分别为：mdIncomingKey 和 mdOutgoingKey 。

服务端的 ctx 和 md 直接打印出来，如下样子：

```
fmt.Println(ctx)
fmt.Println(md)

context.Background.WithValue(type transport.connectionKey, val <not Stringer>).WithValue(type peer.peerKey, val <not Stringer>).WithDeadline(2024-02-19 10:02:43.212614653 +0800 CST m=+41018.106555206 [1.999790196s]).WithValue(type metadata.mdIncomingKey, val <not Stringer>).WithValue(type grpc.streamKey, val <not Stringer>).WithValue(type baggage.baggageContextKeyType, val <not Stringer>).WithValue(type trace.traceContextKeyType, val <not Stringer>).WithValue(type trace.traceContextKeyType, val <not Stringer>).WithCancel

map[:authority:[add.rpc] append:[append-value] content-type:[application/grpc] extra:[extra-value] grpc-accept-encoding:[gzip] noncestr:[abc] signature:[0123456789] timestamp:[2021-07-01 00:00:00] traceparent:[00-89415f99d44e6f8f6e14e3fe8f13ad20-bf33b29c4362ca6a-00] user-agent:[grpc-go/1.59.0]]
signature: [0123456789]
```

注意 md 中的值会被加上中括号“[]”。
