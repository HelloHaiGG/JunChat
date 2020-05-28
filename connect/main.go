package main

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/common/ietcd"
	"JunChat/config"
	connect "JunChat/connect/protocols"
	servers2 "JunChat/connect/servers"
	core "JunChat/core/protocols"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

var RPCServer *string
var NETServer *string

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
	RPCServer = flag.String("RPC", "node-1", "RPC节点")
	NETServer = flag.String("NET", "node-1", "NET节点")
	servers2.NETServer = *NETServer
	flag.Parse()
	log.Println("RPCServer:",*RPCServer,"NETServer:",*NETServer)
	rpcPort, ok := config.APPConfig.CN.Nodes[*RPCServer]
	if !ok {
		log.Fatalf("未配置的RPC节点")
	}
	netPort, ok := config.APPConfig.JC.Nodes[*NETServer]
	if !ok {
		log.Fatalf("未配置的NET节点")
	}

	go servers2.NetConnect(netPort)

	//向Core层汇报
	conn := common.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewCenterServerClient(conn)
	rsp, err := client.OnServerChange(context.Background(), &core.ReportServerStatusParams{ServerId: *RPCServer, Status: common2.NodeStart})
	if err != nil || rsp.Code != common2.Success {
		log.Println("Report Server Status Err:", err)
	}

	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[Connect] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Connect, rpcPort)
	if err != nil {
		log.Fatal("[Connect] Register Server:", err)
	}

	err = register.RunRpcServer(rpcPort, func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers2.ConnectDialServer))
		connect.RegisterPushMsgToConnectServer(server, new(servers2.PushMessageController))
	})

	if err != nil {
		log.Fatal("[Connect] RunRpcServer Err:", err)
	}
}
