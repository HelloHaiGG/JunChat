package common

import (
	"JunChat/common/ietcd"
	"JunChat/config"
	"testing"
	"time"
)

const TestServerName = "TestServer"

func Test_server_register(t *testing.T) {
	config.Init("/Users/mac126/workspace/go-project/HelloMyWorld/config.yaml")
	ietcd.Init(ietcd.IOptions{
		Name:          "",
		Password:      "",
		Hosts:         config.APPConfig.Etcd.Hosts,
		KeepAliveTime: time.Duration(config.APPConfig.Etcd.DialKeepAliveTime),
		DialTimeOut:   time.Duration(config.APPConfig.Etcd.DialTimeOut),
	})
	register, err := NewRegisterSvr(ietcd.Client, 10)
	if err != nil {
		t.Fatal(err)
	}
	err = register.Register(TestServerName, "3838")
	if err != nil {
		t.Fatal(err)
	}

	select {}
}
