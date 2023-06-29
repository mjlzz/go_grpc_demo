# grpc compression
play grpc wiht cumstomized compressor

## Docs

### Encoding

- [grpc encoding](https://github.com/grpc/grpc-go/blob/v1.37.0/Documentation/encoding.md)
	> The gRPC API for sending and receiving is based upon `messages`. However, messages cannot be transmitted directly over a network; they must first be converted into `bytes`. This document describes how gRPC-Go converts messages into bytes and vice-versa for the purposes of network transmission.
    - [encoding api doc](https://pkg.go.dev/google.golang.org/grpc/encoding)
    - [encoding code](https://github.com/grpc/grpc-go/tree/v1.37.0/encoding) define two interface: `Compressor`, `Codec`

When [prepare msg](https://github.com/grpc/grpc-go/blob/v1.37.0/stream.go#L1578), there are two processing steps:
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

### Compress

- [compression Readme](https://github.com/grpc/grpc-go/blob/v1.37.0/Documentation/compression.md)
- [official docs compression](https://grpc.io/docs/guides/compression/)
- [official demo](https://github.com/grpc/grpc-go/tree/v1.37.0/examples/features/compression)
