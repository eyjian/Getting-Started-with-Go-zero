all: user_service

user_service: user.go internal/logic/queryuserlogic.go
	go build -o user_service user.go

clean:
	rm -f user_service

rpc:
	goctl rpc protoc user.proto --go_out=./protoc --go-grpc_out=./protoc --zrpc_out=.

tidy:
	go mod tidy
