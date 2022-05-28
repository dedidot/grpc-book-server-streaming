run-protoc:
	protoc -Iproto --go_out=. --go_opt=module=github.com/dedidot/grpc-book-server-streaming --go-grpc_out=. --go-grpc_opt=module=github.com/dedidot/grpc-book-server-streaming proto/book.proto

build:
	go build -o .