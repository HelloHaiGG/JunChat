package servers

import (
	"JunChat/common"
	"JunChat/common/iredis"
	"context"
)

type MsgListener struct {
	MsgChan chan string
	cxt     context.Context
	cancel  context.CancelFunc
}

func (p *MsgListener) Init() {
	p.MsgChan = make(chan string)
	p.cxt, p.cancel = context.WithCancel(context.Background())
}

func (p *MsgListener) ListenStart() {
	for {
		select {
		case <-p.cxt.Done():
			close(p.MsgChan)
		default:
			results, _ := iredis.RedisCli.BRPop(0, common.MsgQueue).Result()
			p.MsgChan <- results[1]
		}
	}
}

func (p *MsgListener) Stop() {
	p.cancel()
}
