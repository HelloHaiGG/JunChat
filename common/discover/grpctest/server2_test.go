package grpctest

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"time"
)

//测试多个server的负载均衡

//定义结构体,实现服务端方法
type HServer struct{}

func (p *HServer) SayHello(ctx context.Context, in *Request) (*Response, error) {
	fmt.Println("server 2 be call...")
	return &Response{Server: "Hello Server 2", Msg: in.Msg}, nil
}

func TestRegisterHelloServer2(t *testing.T) {
	config.Init("/Users/mac126/workspace/go-project/HelloMyWorld/config.yaml")
	ietcd.Init(ietcd.IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})
	register, err := serverhandler.NewRegisterSvr(ietcd.Client, 10)
	if err != nil {
		t.Fatal(err)
	}

	err = register.Register(testServer, "3839")
	if err != nil {
		t.Fatal(err)
	}

	err = register.RunRpcServer("3839", func(s *grpc.Server) {
		RegisterHelloServer(s, new(HServer))
	})
	if err != nil {
		t.Fatal(err)
	}
}
