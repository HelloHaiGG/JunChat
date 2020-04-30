package servers

import (
	"JunChat/queue/modles"
	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/common/log"
)

var Listener MsgListener

func Start() {
	Listener.Init()
	go Listener.ListenStart()
	go MsgDistribute()
}

func MsgDistribute() {
	for {
		select {
		case msg, ok := <-Listener.MsgChan:
			if !ok {
				return
			}
			body := &modles.MsgWrap{}
			_ = jsoniter.UnmarshalFromString(msg, body)
			log.Info("Lop Msg In Queue Suc,Msg Id:",body.Id)
			body.PushMsg()
		}
	}
}
