#### Streaming server gRPC

#### How to run ?
```bash
# Run protoc
make run-protoc

# Build app
make build

# Run app
./grpc-book-server-streaming
```

#### API Endpoint : http://0.0.0.0:12345
```
// Process
main.go -> server.go -> service.go -> repository.go
```

#### Run test
https://medium.com/easyread/testing-your-grpc-using-postman-5167347f26dd