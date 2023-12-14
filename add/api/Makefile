all: add_http_server
add_http_server: adder.go internal/logic/addlogic.go
	go build -o add_http_server adder.go

api: add.api
	goctl api go -api add.api -dir .
	go mod tidy

clean:
	rm -f add_http_server
