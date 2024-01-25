go-zero 的网关往 rpc 服务传递数据时，可以使用 headers，但需要注意前缀规则，否则会发现数据传递不过去，或者对方取不到数据。

go-zero 的网关对服务的调用使用了第三方库 grpcurl，入口函数为 InvokeRPC：

```go
grpcurl.InvokeRPC(r.Context(), source, cli.Conn(), rpcPath, s.prepareMetadata(r.Header), handler, parser.Next)
```

调用在 https://github.com/zeromicro/go-zero/blob/master/gateway/server.go 中进行的，上述调用会处理 HTTP 的 headers 数据，对于不是以字符串“Grpc-Metadata-”打头的会过滤掉，对于以字符串“Grpc-Metadata-”打头的会将“Grpc-Metadata-”转为“gateway-”。

```go
// go-zero/gateway/internal/headerprocessor.go
// ProcessHeaders builds the headers for the gateway from HTTP headers.
func ProcessHeaders(header http.Header) []string {
	var headers []string

	for k, v := range header {
		if !strings.HasPrefix(k, metadataHeaderPrefix) { // 判断是否包含了前缀“Grpc-Metadata-”
			continue // 如果没有前缀“Grpc-Metadata-”则直接过滤丢弃掉
		}

		// 将前缀“Grpc-Metadata-”替换为前缀“gateway-”
		key := fmt.Sprintf("%s%s", metadataPrefix, strings.TrimPrefix(k, metadataHeaderPrefix))
		for _, vv := range v {
			headers = append(headers, key+":"+vv)
		}
	}

	return headers
}
```

函数 MetadataFromHeaders 负责从 headers 解码数据：

```
// https://github.com/fullstorydev/grpcurl/blob/master/grpcurl.go
func MetadataFromHeaders(headers []string) metadata.MD {
	md := make(metadata.MD)
	for _, part := range headers {
		if part != "" {
			pieces := strings.SplitN(part, ":", 2)
			if len(pieces) == 1 {
				pieces = append(pieces, "") // if no value was specified, just make it "" (maybe the header value doesn't matter)
			}
			headerName := strings.ToLower(strings.TrimSpace(pieces[0]))
			val := strings.TrimSpace(pieces[1])
			if strings.HasSuffix(headerName, "-bin") {
				if v, err := decode(val); err == nil {
					val = v
				}
			}
			md[headerName] = append(md[headerName], val)
		}
	}
	return md
}
```

```go
// https://github.com/fullstorydev/grpcurl/blob/master/invoke.go
func InvokeRPC(ctx context.Context, source DescriptorSource, ch grpcdynamic.Channel, methodName string,
	headers []string, handler InvocationEventHandler, requestData RequestSupplier) error {

	md := MetadataFromHeaders(headers)

	svc, mth := parseSymbol(methodName)
	if svc == "" || mth == "" {
		return fmt.Errorf("given method name %q is not in expected format: 'service/method' or 'service.method'", methodName)
	}

	dsc, err := source.FindSymbol(svc)
	if err != nil {
		// return a gRPC status error if hasStatus is true
		errStatus, hasStatus := status.FromError(err)
		switch {
		case hasStatus && isNotFoundError(err):
			return status.Errorf(errStatus.Code(), "target server does not expose service %q: %s", svc, errStatus.Message())
		case hasStatus:
			return status.Errorf(errStatus.Code(), "failed to query for service descriptor %q: %s", svc, errStatus.Message())
		case isNotFoundError(err):
			return fmt.Errorf("target server does not expose service %q", svc)
		}
		return fmt.Errorf("failed to query for service descriptor %q: %v", svc, err)
	}
	sd, ok := dsc.(*desc.ServiceDescriptor)
	if !ok {
		return fmt.Errorf("target server does not expose service %q", svc)
	}
	mtd := sd.FindMethodByName(mth)
	if mtd == nil {
		return fmt.Errorf("service %q does not include a method named %q", svc, mth)
	}

	handler.OnResolveMethod(mtd)

	// we also download any applicable extensions so we can provide full support for parsing user-provided data
	var ext dynamic.ExtensionRegistry
	alreadyFetched := map[string]bool{}
	if err = fetchAllExtensions(source, &ext, mtd.GetInputType(), alreadyFetched); err != nil {
		return fmt.Errorf("error resolving server extensions for message %s: %v", mtd.GetInputType().GetFullyQualifiedName(), err)
	}
	if err = fetchAllExtensions(source, &ext, mtd.GetOutputType(), alreadyFetched); err != nil {
		return fmt.Errorf("error resolving server extensions for message %s: %v", mtd.GetOutputType().GetFullyQualifiedName(), err)
	}

	msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(&ext)
	req := msgFactory.NewMessage(mtd.GetInputType())

	handler.OnSendHeaders(md)
	ctx = metadata.NewOutgoingContext(ctx, md)

	stub := grpcdynamic.NewStubWithMessageFactory(ch, msgFactory)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if mtd.IsClientStreaming() && mtd.IsServerStreaming() {
		return invokeBidi(ctx, stub, mtd, handler, requestData, req)
	} else if mtd.IsClientStreaming() {
		return invokeClientStream(ctx, stub, mtd, handler, requestData, req)
	} else if mtd.IsServerStreaming() {
		return invokeServerStream(ctx, stub, mtd, handler, requestData, req)
	} else {
		return invokeUnary(ctx, stub, mtd, handler, requestData, req)
	}
}
```

网关可如下实现：

```go
newReq := r.WithContext(r.Context())
newReq.Header.Set("Grpc-Metadata-myuid", userId)
next.ServeHTTP(w, newReq)
```

服务端的实现：

```go
vals := metadata.ValueFromIncomingContext(l.ctx, "gateway-myuid")
userId := vals[0]
```
