# protobuf

protobuf is just a serialization/deserialization tool (just like JSON)


## Question

what will happen if proto files mismatch?

- rpc method name
- rpc request name
- rpc request parameter name
- rpc request parameter type


## Demo

1. start server

- server use ./proto

```bash
go run server/server.go
```


2. send request

- client use ./proto2

```bash
go run client/client.go
```
