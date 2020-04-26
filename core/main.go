package main

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/common/ietcd"
	"JunChat/common/igorm"
	"JunChat/common/iredis"
	"JunChat/config"
	"JunChat/core/models"
	core "JunChat/core/protocols"
	"JunChat/core/servers"
	"JunChat/utils"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
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
	igorm.Init(&igorm.IOption{
		Host:     "",
		Port:     0,
		User:     "root",
		Password: "starunion",
		DB:       "jun_chat",
		IsDebug:  true,
	})

	igorm.DbClient.AutoMigrate(&models.UserInfo{})

	iredis.Init(&iredis.IOptions{DialTimeOut:10 * time.Second,MaxConnAge:10*time.Second})

	//通过命令控制运行端口
	ports := make(map[string]string)
	for i, node := range config.APPConfig.CC.Nodes {
		ports[fmt.Sprintf("node-%d", i+1)] = node
	}
	NODE := flag.String("node", "node-1", "程序节点")
	flag.Parse()
	port, ok := ports[*NODE]
	if !ok {
		log.Fatalf("未配置的节点")
	}


	//初始化调度Map
	servers.InitDispatchMap()
	//初始化雪花生成器
	machineId, _ := strconv.ParseInt(port, 10, 64)
	utils.InitSF(machineId%10)
	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[Core] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Core, port)
	if err != nil {
		log.Fatal("[Core] Register Server:", err)
	}

	err = register.RunRpcServer(port, func(server *grpc.Server) {
		common2.RegisterProtoDialServer(server, new(servers.CoreDialServer))
		core.RegisterUserControllerServer(server, new(servers.UserController))
		core.RegisterSendMsgControllerServer(server, new(servers.SendMessageController))
		core.RegisterCenterServerServer(server, new(servers.CenterServerController))
	})

	if err != nil {
		log.Fatal("[Core] RunRpcServer Err:", err)
	}
}
