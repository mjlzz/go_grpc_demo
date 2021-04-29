# protobuf

protobuf is just a serialization/deserialization tool (just like JSON)

[Protobuf语法](https://segmentfault.com/a/1190000007917576)


## Question

what will happen if proto files mismatch?

- rpc method name
- rpc request name
- rpc request parameter name
- rpc request parameter type

./proto/play.proto
```protobuf
...

service PlaySvc {
    rpc Play(PlayReq) returns (PlayRes){}
    rpc Stop(StopReq) returns (StopRes){}
}

message PlayReq {
    string Address = 1;
    int64 Nums = 2;
}

...
```

./proto2/play.proto
```protobuf
...

service PlaySvc {
    rpc Play(PlayReq2) returns (PlayRes){}
    rpc Stop2(StopReq) returns (StopRes){}
}

message PlayReq2 {
    string A = 1;
    string Cat = 2;
    int64  Nums = 3;
}

...
```

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
