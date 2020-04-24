package ietcd

import (
	"testing"
	"time"
)

func TestEtcdInit(t *testing.T) {
	config.Init("/Users/mac126/workspace/go-project/HelloMyWorld/config.yaml")
	Init(IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})

}
