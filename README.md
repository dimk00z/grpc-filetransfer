# grpc-filetransfer

## Article

https://dev.to/dimk00z/grpc-file-transfer-with-go-1nb2

## Server 

```go run cmd/server/main.go```

Server config `./config/server/config.yml` or you can use envs.

## Client

```
go run cmd/client/main.go -h
Sending files via gRPC

Usage:
  transfer_client [flags]

Flags:
  -a, --addr string   server address
  -b, --batch int     batch size for sending (default 1048576)
  -f, --file string   file path
  -h, --help          help for transfer_client
```

File transfer speed depends on `batch` size 

### Client run

```go run cmd/client/main.go -a=':9000' -f=8GB.bin```
