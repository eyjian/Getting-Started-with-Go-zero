all: add_client add_server

rpc:
	goctl rpc protoc add.proto --go_out=./protoc --go-grpc_out=./protoc --zrpc_out=.
	go mod tidy

add_client: add_client.go
	go build -o add_client add_client.go

add_server: add.go internal/logic/addlogic.go
	go build -o add_server add.go

clean:
	rm -f add_server add_client
