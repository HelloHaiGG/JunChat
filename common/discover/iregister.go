package common

import (
	"JunChat/common/ietcd"
	"JunChat/utils"
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"net"
	"time"
)

type RegisterSvr struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	cancelLease   context.CancelFunc
}

func NewRegisterSvr(client *clientv3.Client, ttl int64) (*RegisterSvr, error) {

	register := new(RegisterSvr)

	if client == nil {
		log.Fatal("Etcd Client is Nil. First 'ietcd.Init()'")
	}

	register.client = ietcd.Client

	if err := register.keepAlive(ttl); err != nil {
		return nil, err
	}

	return register, nil
}

//续租
func (p *RegisterSvr) keepAlive(ttl int64) error {
	lease := clientv3.NewLease(p.client)
	p.lease = lease

	//以包活时间为超时时间
	outCxt, _ := context.WithTimeout(context.Background(), time.Duration(ttl)*time.Second)
	resp, err := lease.Grant(outCxt, ttl)
	p.leaseResp = resp
	if err != nil {
		return err
	}
	cxt, cancel := context.WithCancel(context.TODO())
	p.cancelLease = cancel
	//继租
	p.keepAliveChan, err = lease.KeepAlive(cxt, resp.ID)
	if err != nil {
		return err
	}

	//监听租约情况
	go p.ListenKeepAliveResp()

	return nil
}

//监听续租情况
func (p *RegisterSvr) ListenKeepAliveResp() {
	for {
		select {
		case resp := <-p.keepAliveChan:
			if resp == nil {
				log.Println("Etcd 续约失败.")
				time.Sleep(10 * time.Second)
			}
		}
	}
}

//撤销租约
func (p *RegisterSvr) CancelLease() error {
	p.cancelLease()
	_, err := p.lease.Revoke(context.TODO(), p.leaseResp.ID)
	return err
}

//注册服务
func (p *RegisterSvr) Register(server string, port string) error {
	kv := clientv3.NewKV(p.client)
	addr := net.JoinHostPort(utils.GetInternalIp(), port)
	server = fmt.Sprintf("/%s/%s/%s", DNSName, server, addr)
	_, err := kv.Put(context.TODO(), server, addr, clientv3.WithLease(p.leaseResp.ID))

	return err
}
