# Stream



## Case
> 参考 [文档1](https://cloud.tencent.com/developer/article/1445070), [文档2](https://zhuanlan.zhihu.com/p/141677241)

- service-side rpc streaming 场景：
  - 股票app：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端
  - app的在线push：client先发请求到server注册，然后server就可以发在线push了

- client-side rpc streaming 场景：
  - 数据上传(上万条记录)：如果只用simple rpc的话，就要一次性收到上万条记录，并且在这些记录传输完成之后才能对数据进行处理。如果用streaming rpc的话，可以在收到一些记录之后就开始处理，以此减少了服务器的瞬时压力，也更有实时性
  - 客户端并发调用细小粒度的接口。比如有5个后台接口A B C D E，客户端在不同页面，可以调用不同的接口组合。比如在个人页，就调用ABC；在动态页面，就调用CDE，后台都只会有一个rsp。这种模式的好处就是让后台可以将接口的粒度细化，客户端调用灵活，减少重复代码，提高复用率

- bi-side rpc streaming （类似于tcp）场景：
  - 聊天机器人
  - 有状态的游戏服务器进行数据交换。比如LOL，王者荣耀等竞技游戏，client和server之间需要非常频繁地交换数据
