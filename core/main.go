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
		log.Fatal("[Core] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Core, "4953")
	if err != nil {
		log.Fatal("[Core] Register Server:", err)
	}

	err = register.RunRpcServer("4953", func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers.CoreDialServer))
	})

	if err != nil {
		log.Fatal("[Core] RunRpcServer Err:", err)
	}
}
