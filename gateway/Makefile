all: gateway
gateway: main.go
	go build -o gateway main.go 

clean:
	rm -f gateway

rpc:
	goctl gateway -dir gateway

