package common

import (
	"google.golang.org/grpc/resolver"
	"log"
)

/**
基于 GRPC 提供的 resolver 接口实现服务发现与注册,并提供了负载均衡策略
IBuilder 创建一个解析器
*/

func init() {
	//注册解析器
	resolver.Register(NewIBuilder())
}

const DNSName  = "JXian-Server"

// IBuilder实现 实现Builder接口
type IBuilder struct {
	DNS string
}

func NewIBuilder() *IBuilder {
	return &IBuilder{
		DNS:     DNSName,
	}
}

// GRPC 将所有配置的 Resolver 维护到一个ResolverMap  --> map[string]Resolver
// GRPC 通过 Scheme() 获取 Resolver 的 key
func (p *IBuilder) Scheme() string {
	return p.DNS
}

func (p *IBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	if ietcd.Client == nil{
		log.Fatal("Etcd Client is Nil. First 'ietcd.Init()'")
	}

	//创建 IResolver
	iResolver := &IResolver{
		cc: cc,
		c:  ietcd.Client,
	}

	go iResolver.watch("/"+target.Scheme + "/" +target.Endpoint)

	return iResolver,nil
}






