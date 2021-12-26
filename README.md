# go grpc demo
Different grpc usages implemented by golang

## Demo list

- [multiple proto](multi_proto#readme)
  - register multiple pb servers on the same grpc server

- [codec](codec#readme)
  - play grpc wiht cumstomized codec

- [mismatch proto](proto_mismatch#readme)
  - what will happen if client's proto mismatch with server's proto?

- [health check](health_check#readme)
  - monitor connetion status by strobing server

- [stream usages](stream#readme)
  - one unary usage and three different stream usages

- [go kit](go_kit#readme)
  - a popular kit for golang, play it with grpc


## Before running demo
- Init project

  - [update go proxy](https://goproxy.io/zh/)

  - [download all the dependencies](https://golangbyexample.com/go-mod-tidy/)

- Compile proto

  - [compile cmd](go_kit/pb/compile.sh)


## Docs
- [Go Quick start](https://grpc.io/docs/languages/go/quickstart/)
- [gRPC-Go API](https://pkg.go.dev/google.golang.org/grpc)
