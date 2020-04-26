package main

import (
	"JunChat/common/ietcd"
	"JunChat/config"
	"JunChat/gateway/router"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
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

	eg := &errgroup.Group{}

	//用户网关服务
	junServer := &http.Server{
		Addr:         "127.0.0.1:59277", //及鲜app
		Handler:      router.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	eg.Go(func() error {
		return junServer.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		log.Fatal("App gateway err:", err)
	}
}
