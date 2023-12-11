all: gateway_login

gateway_login: main.go
	go build -o gateway_login main.go

clean:
	rm -f gateway_login

rpc:
	goctl gateway -dir gateway

pb:
	protoc --descriptor_set_out=proto/login.pb proto/login.proto
	protoc --descriptor_set_out=proto/user.pb proto/user.proto
