# go_pubsub

go_pubsub 是基于redis协议实现的pubsub服务端. 由于兼容redis protocol, 可直接使用redis client及sdk操作.

`to do list:`

* 持久化
* 丰富增删改查

## 演示:

`server:`

```
go run cmd/main.go
```

`client:`

```
[gopy@xiaorui ~ ]$ redis-cli -p 9999
127.0.0.1:9999> PUBLISH xiaorui.cc hello
OK
127.0.0.1:9999> SUBSCRIBE xiaorui.cc
Reading messages... (press Ctrl-C to quit)
hello
```
