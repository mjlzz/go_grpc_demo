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
    int64 Nums = 1;
    string Address = 2;
    int64 Count = 3;
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

## Mismatch

sender encode reqeust/respose struct to protobuf bytes, dismiss attribute's name, send attribute with order number and type

transport bytes

receiver decode protobuf bytes by order and type, if types mismatch in the same order, save bytes into `unknownFields` and set corresponding attribute default value
