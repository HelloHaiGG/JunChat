package servers

import (
	"JunChat/common"
	"JunChat/queue/modles"
	jsoniter "github.com/json-iterator/go"
)

var Listener MsgListener

func init() {
	Listener.Init()
	Listener.ListenStart()
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
			if body.MsgType == common.SingleChat {
				body.PushMsg()
				continue
			}
		}
	}
}
