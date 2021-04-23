# go-kit with grpc

[go-kit doc](https://gokit.io/examples/)

[go-kit api](https://pkg.go.dev/github.com/go-kit/kit@v0.10.0)

[example](https://github.com/go-kit/kit/tree/master/examples/addsvc)

## demo

1. start server

```bash
go run cmd/svr/svc.go
```

2. client request

```bash
go run cmd/cli/addcli.go -grpc-addr :8082 -method sum 1 4
```
