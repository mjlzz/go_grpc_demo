# grpc codec

## Docs

- [grpc encoding](https://github.com/grpc/grpc-go/blob/v1.37.0/Documentation/encoding.md)
> The gRPC API for sending and receiving is based upon `messages`. However, messages cannot be transmitted directly over a network; they must first be converted into `bytes`. This document describes how gRPC-Go converts messages into bytes and vice-versa for the purposes of network transmission.
    - [encoding api doc](https://pkg.go.dev/google.golang.org/grpc/encoding)
    - [encoding code](https://github.com/grpc/grpc-go/tree/v1.37.0/encoding)


- [Codec interface](https://github.com/grpc/grpc-go/blob/v1.37.0/encoding/encoding.go#L86)
```go
// Codec defines the interface gRPC uses to encode and decode messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
}
```

- protobuf codec
    - [proto api doc](https://pkg.go.dev/google.golang.org/grpc/encoding/proto)
    - [proto code](https://github.com/grpc/grpc-go/tree/v1.37.0/encoding/proto)


## Demo

- Define customized codec function based on [Codec interface](https://github.com/grpc/grpc-go/blob/v1.37.0/encoding/encoding.go#L86)
    - [encoding/json](https://pkg.go.dev/encoding/json)

1. start server
```bash
go run ./greeter_server/main.go
```

2. send request
```bash
go run ./greeter_client/main.go
```
