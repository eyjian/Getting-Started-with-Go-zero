all: authentication_service

authentication_service: authentication.go
	go build -o authentication_service authentication.go

rpc:
	goctl rpc protoc authentication.proto --go_out=./protoc --go-grpc_out=./protoc --zrpc_out=.

# 在 rpc 后执行
tidy:
	go mod tidy

clean:
	rm -f authentication_service
