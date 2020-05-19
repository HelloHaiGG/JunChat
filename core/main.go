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
		Host:     config.APPConfig.Mysql.Host,
		Port:     config.APPConfig.Mysql.Port,
		User:     config.APPConfig.Mysql.User,
		Password: config.APPConfig.Mysql.Password,
		DB:       "jun_chat",
		IsDebug:  true,
	})

	igorm.DbClient.AutoMigrate(&models.UserInfo{})

	iredis.Init(&iredis.IOptions{
		Host:config.APPConfig.Redis.Host,
		Port:config.APPConfig.Redis.Port,
		Password:config.APPConfig.Redis.Password,
		DB:config.APPConfig.Redis.DB,
		DialTimeOut: 10 * time.Second,
		MaxConnAge: 10 * time.Second,
	})

	//初始化缓存信息 【玩家Token,服务负载信息】
	//iredis.RedisCli.Del("")

	//通过命令控制运行端口
	NODE := flag.String("RPC", "core-1", "程序节点")
	flag.Parse()
	port, ok := config.APPConfig.CC.Nodes[*NODE]
	if !ok {
		log.Fatalf("未配置的节点")
	}

	//初始化调度Map
	servers.InitDispatchMap()
	//初始化雪花生成器
	machineId, _ := strconv.ParseInt(port, 10, 64)
	utils.InitSF(machineId % 10)
	//注册服务
	register, err := common.NewRegisterSvr(ietcd.Client, int64(config.APPConfig.Grpc.CallTimeOut))
	if err != nil {
		log.Fatal("[Core] New Register Server:", err)
	}
	err = register.Register(config.APPConfig.Servers.Core, port)
	if err != nil {
		log.Fatal("[Core] Register Server:", err)
	}
	//注册RPC服务
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
