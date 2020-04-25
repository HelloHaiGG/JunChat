package main

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/common/ietcd"
	"JunChat/config"
	"JunChat/core/servers"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	config.Init("./config.yaml")
	ietcd.Init(ietcd.IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})

	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[GateWay] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Gateway, "5164")
	if err != nil {
		log.Fatal("[GateWay] Register Server:", err)
	}

	err = register.RunRpcServer("5164", func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers.CoreDialServer))
	})

	if err != nil {
		log.Fatal("[GateWay] RunRpcServer Err:", err)
	}
}