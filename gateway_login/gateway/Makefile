all: gateway_login

# 执行次序：
# make rpc
# make pb
# make tidy
# make

gateway_login: main.go middleware/error.go middleware/login.go middleware/auth.go
	go build -o gateway_login main.go config.go

tidy:
	go mod tidy

clean:
	rm -f gateway_login

rpc:
	goctl gateway -dir gateway
	go mod tidy

pb:
	protoc --descriptor_set_out=proto/login.pb proto/login.proto
	protoc --descriptor_set_out=proto/user.pb proto/user.proto
