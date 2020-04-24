package grpctest

import (
	"HelloMyWorld/common/ietcd"
	"HelloMyWorld/common/serverholder"
	"HelloMyWorld/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"time"
)

// grpc-etcd 服务端测试

//定义结构体,实现服务端方法
type HaServer struct{}

func (p *HaServer) SayHello(ctx context.Context, in *Request) (*Response, error) {
	fmt.Println("server be call...")
	return &Response{Server: "Hello Server", Msg: in.Msg}, nil
}

const testServer = "TestServer"

func TestRegisterHelloServer(t *testing.T) {
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

	err = register.Register(testServer, "3838")
	if err != nil {
		t.Fatal(err)
	}

	err = register.RunRpcServer("3838", func(s *grpc.Server) {
		RegisterHelloServer(s, new(HaServer))
	})
	if err != nil {
		t.Fatal(err)
	}
}
