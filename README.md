# JunChat
Go-分布式聊天系统

#### 结构
![png](./JunChat.png)


#### 流程

````
1. 网关层 
    提供用户登录,注册,发送消息接口
2. 链接层（可启动多个server）
    1).维护每一个用户的链接,以及聊天室
    2).链接断开后向中心逻辑层汇报
3. 中心逻辑层 (可启动多个server)
    1).接收网关层RPC,处理玩家注册,登录逻辑
    2).通过负载均衡,对玩家分配 Conncet Server
    3).维护玩家所在Server,并将数据持久化到Reids
    4).接收网关层RPC调用,处理消息发送
    5).处理(解密,解压等)好消息后将消息推送到Redis队列,供Queue层消费
    6).接收链接层汇报,将同步玩家退出的消息
4. 队列层
    提供消息队列,接受逻辑层处理好的消息,并将消息发送到对应的Connect Server
````
````
Gateway:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway main.go
Connect:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o connect main.go
Core:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o core main.go
Queue:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o queue  main.go
````