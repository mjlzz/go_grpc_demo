# grpc codec
play grpc wiht cumstomized codec


## Docs

### Encoding

- [grpc encoding](https://github.com/grpc/grpc-go/blob/v1.37.0/Documentation/encoding.md)
	> The gRPC API for sending and receiving is based upon `messages`. However, messages cannot be transmitted directly over a network; they must first be converted into `bytes`. This document describes how gRPC-Go converts messages into bytes and vice-versa for the purposes of network transmission.
    - [encoding api doc](https://pkg.go.dev/google.golang.org/grpc/encoding)
    - [encoding code](https://github.com/grpc/grpc-go/tree/v1.37.0/encoding) define two interface: `Compressor`, `Codec`

When [prepare msg](https://github.com/grpc/grpc-go/blob/v1.37.0/stream.go#L1578), there are two encoding steps:
      1. [encode](https://github.com/grpc/grpc-go/blob/v1.37.x/rpc_util.go#L588)
      2. [compress](https://github.com/grpc/grpc-go/blob/v1.37.x/rpc_util.go#L606)
  	. If not set compress type, `msgHeader` will use raw data.
```go
	// The input interface is not a prepared msg.
	// Marshal and Compress the data at this point
	data, err = encode(codec, m)
	if err != nil {
		return nil, nil, nil, err
	}
	compData, err := compress(data, cp, comp)
	if err != nil {
		return nil, nil, nil, err
	}
	hdr, payload = msgHeader(data, compData)
```

### Codec

- grpc register the Codec for "proto" by default
    - [default](https://github.com/grpc/grpc-go/blob/v1.37.0/codec.go#L23)

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

- protobuf codec (one way to codec)
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
