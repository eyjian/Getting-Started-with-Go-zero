all: login_service

login_service: login.go internal/logic/loginlogic.go
	go build -o login_service login.go

clean:
	rm -f login_service

rpc:
	goctl rpc protoc login.proto --go_out=./protoc --go-grpc_out=./protoc --zrpc_out=.

tidy:
	go mod tidy
