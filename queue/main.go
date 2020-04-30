package main

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/common/ietcd"
	"JunChat/common/iredis"
	"JunChat/config"
	"JunChat/queue/servers"
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

	iredis.Init(&iredis.IOptions{
		Host:        config.APPConfig.Redis.Host,
		Port:        config.APPConfig.Redis.Port,
		Password:    config.APPConfig.Redis.Password,
		DB:          config.APPConfig.Redis.DB,
		DialTimeOut: 10 * time.Second,
		MaxConnAge:  10 * time.Second,
	})

	go servers.Start()
	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[Queue] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Queue, "28484")
	if err != nil {
		log.Fatal("[Queue] Register Server:", err)
	}

	err = register.RunRpcServer("28484", func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers.QueueDialServer))
	})

	if err != nil {
		log.Fatal("[Queue] RunRpcServer Err:", err)
	}
}
