all: gateway_auth
gateway: main.go
	go build -o gateway_auth main.go 

clean:
	rm -f gateway_auth

rpc:
	goctl gateway -dir gateway

