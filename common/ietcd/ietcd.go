package ietcd

import (
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var Client *clientv3.Client

func Init(option IOptions) {
	option.init()

	cfg := clientv3.Config{
		Endpoints:         option.Hosts,
		DialTimeout:       option.DialTimeOut,
		DialKeepAliveTime: option.KeepAliveTime,
		Username:          option.Name,
		Password:          option.Password,
	}
	c, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("Etcd New Client Err : %v \n", err)
	}
	Client = c
}

type IOptions struct {
	Name          string
	Password      string
	Hosts         []string
	KeepAliveTime time.Duration
	DialTimeOut   time.Duration
}

func (p *IOptions) init() {
	if p.Password == "" {
		p.Password = ""
	}
	if p.Name == "" {
		p.Name = "root"
	}
	if p.Hosts == nil || len(p.Hosts) == 0 {
		p.Hosts = []string{"127.0.0.1:2379"}
	}
	if p.KeepAliveTime == 0 {
		p.KeepAliveTime = time.Second * 10
	}
	if p.DialTimeOut == 0 {
		p.DialTimeOut = time.Second * 5
	}
}
