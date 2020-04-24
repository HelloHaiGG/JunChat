package common

import (
	"HelloMyWorld/config"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"sync"
	"time"
)

/**
GRPC的ClientConn对象可以帮我们实现自动重连的机制，并且是并发安全的，因此可以定义一个全局的ClientConn。
*/

var global sync.Map

func GetServerConn(server string) *grpc.ClientConn {
	if value, ok := global.Load(server); !ok {
		return NewRpcConn(server)
	} else {
		return value.(*grpc.ClientConn)
	}
}
func NewRpcConn(server string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		//grpc.WithBlock(),
	}
	cxt, _ := context.WithTimeout(context.Background(), time.Duration(config.APPConfig.Grpc.CallTimeOut)*time.Second)
	conn, err := grpc.DialContext(cxt, DNSName+":///"+server, opts...)
	if err != nil {
		log.Fatalf("Grpc Dial %s Error:%v", server, err)
	}
	global.Store(server, conn)
	return conn
}
