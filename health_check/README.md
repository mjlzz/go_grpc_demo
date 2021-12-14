# grpc_health_check

## Docs

[grpc doc](https://pkg.go.dev/google.golang.org/grpc)

[grpc health doc](https://pkg.go.dev/google.golang.org/grpc@v1.37.0/health)

[health source code](https://sourcegraph.com/github.com/grpc/grpc-go/-/tree/health)

[tutorial](https://github.com/grpc-ecosystem/grpc-health-probe)

[grpc health-checking protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)


## Demo
0. init project

- [update go proxy](https://goproxy.io/zh/)

- [download all the dependencies](https://golangbyexample.com/go-mod-tidy/)

1. start server

```bash
go run ./server/server.go
```

2. in another terminal

```bash
./bin/grpc-health-probe -addr 127.0.0.1:50051
```
