package main

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/common/ietcd"
	"JunChat/config"
	connect "JunChat/connect/protocols"
	servers2 "JunChat/connect/servers"
	"JunChat/queue/servers"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

var NODE *string

func main() {
	config.Init("./config.yaml")
	ietcd.Init(ietcd.IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})
	//通过命令控制运行端口
	ports := make(map[string]string)
	netPorts := make(map[string]string)
	for i, node := range config.APPConfig.CN.Nodes {
		key := fmt.Sprintf("node-%d", i+1)
		ports[key] = node
		netPorts[key] = config.APPConfig.JC.Nodes[i]
	}
	NODE = flag.String("node", "node-1", "程序节点")
	servers2.NODE = *NODE
	flag.Parse()
	port, ok := ports[*NODE]
	if !ok {
		log.Fatalf("未配置的节点")
	}

	go servers2.NetConnect(netPorts[*NODE])

	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[Connect] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Connect, port)
	if err != nil {
		log.Fatal("[Connect] Register Server:", err)
	}

	err = register.RunRpcServer(port, func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers.QueueDialServer))
		connect.RegisterPushMsgToConnectServer(server, new(servers2.PushMessageController))
	})

	if err != nil {
		log.Fatal("[Connect] RunRpcServer Err:", err)
	}
}
