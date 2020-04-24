package common

import (
	"context"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

/**
实现 Resolver 对象
Resolver:服务地址维护对象,通过两个接口方法可以改变维护状态
*/

const (
	NOT_EXIST = -1
)

type IResolver struct {
	cc   resolver.ClientConn
	c    *clientv3.Client
	addr []resolver.Address
}

func (p *IResolver) ResolveNow(opt resolver.ResolveNowOptions) {

}

func (p *IResolver) Close() {

}

//监控服务地址变化
func (p *IResolver) watch(prefix string) {
	p.GetAddr(prefix)
	p.UpdateAddr(prefix)
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			p.UpdateAddr(prefix)
		}
	}
}

//监控地址的变化
func (p *IResolver) UpdateAddr(prefix string) {
	//prefix 通过服务注册的前缀,可以获取到所有的服务地址  (如: /ISERVER/OrderServer/*** : 127.0.0.1:8888,127.0.0.2:8888)
	state := resolver.State{}
	wChan := p.c.Watch(context.TODO(), prefix, clientv3.WithPrefix())
	for response := range wChan {
		for _, v := range response.Events {
			switch v.Type {
			case mvccpb.PUT:
				if p.isExist(string(v.Kv.Value)) == NOT_EXIST {
					p.addr = append(p.addr, resolver.Address{Addr: string(v.Kv.Value)})
				}
			case mvccpb.DELETE:
				if res := p.isExist(string(v.Kv.Value)); res != NOT_EXIST {
					if len(p.addr) == 0 {
						continue
					}
					p.addr = append(p.addr[:res], p.addr[res+1:]...)
				}
			}
		}
		state.Addresses = p.addr
		p.cc.UpdateState(state)
	}
}

//服务启动后，先将地址加载到 Resolver
func (p *IResolver) GetAddr(prefix string) {
	resp, err := p.c.Get(context.TODO(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Fatal("Get Server Addr error:", err)
	}
	state := resolver.State{}
	for _, kv := range resp.Kvs {
		if p.isExist(string(kv.Value)) == NOT_EXIST {
			p.addr = append(p.addr, resolver.Address{Addr: string(kv.Value)})
		}
	}
	state.Addresses = p.addr
	p.cc.UpdateState(state)
}

func (p *IResolver) isExist(addr string) int32 {
	for k, _ := range p.addr {
		if p.addr[k].Addr == addr {
			return int32(k)
		}
	}
	return NOT_EXIST
}
