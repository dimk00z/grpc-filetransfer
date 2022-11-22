# grpc-filetransfer

## Server 

```go run cmd/server/main.go```

Настройки сервера в `./config/server/config.yml` или можно прокинуть через переменные окружения.

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

Скорость передачи/записи зависит от размера `batch`, нужно подбирать оптимальный 

### Пример запуска клиента 

```go run cmd/client/main.go -a=':9000' -f=8GB.bin```


## Testing

Для оптимальной проверки тут нужны интеграционные тесты, при необходимости подумаю про это.
