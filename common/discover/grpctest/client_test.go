package grpctest

import (
	serverhandler "common/discover"
	"context"
	"fmt"
	"testing"
	"time"
)

// grpc-etcd 客户端测试

func TestGetServerInstance(t *testing.T) {

	config.Init("/Users/mac126/workspace/go-project/HelloMyWorld/config.yaml")
	ietcd.Init(ietcd.IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})

	conn := serverhandler.GetServerConn(testServer)

	client := NewHelloClient(conn)
	for true {
		res, err := client.SayHello(context.Background(), &Request{Msg: "Hello Server !"})
		if err != nil {
			t.Log(err)
		}else{
			fmt.Println("server:", res.Server, " Msg:", res.Msg)
		}
	}

}
