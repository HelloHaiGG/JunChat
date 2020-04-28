package modles

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	connect "JunChat/connect/protocols"
	queue "JunChat/queue/protocols"
	"context"
	"github.com/gogo/protobuf/proto"
	"log"
)

type Node struct {
	Host string `json:"Host,omitempty"`
	Port string `json:"Port,omitempty"`
}

type MsgWrap struct {
	*queue.MessageBody
	N map[string]*Node `json:"N"` //接受者所在节点信息
}

func (p *MsgWrap) PushMsg() {
	for id, node := range p.N {
		conn := common.GetServerConnByHost(node.Host, node.Port)
		client := connect.NewPushMsgToConnectClient(conn)
		b, _ := proto.Marshal(p.MessageBody)
		msg := &connect.MessageBody{}
		_ = proto.Unmarshal(b, msg)
		//将接受者换层UserId
		//聊天室消息在发送时,ReceiverId 是聊天室的Id,现在需要将Id替换成UserId
		msg.Receiver = id
		rsp, err := client.PushMsgToConnectServer(context.Background(), &connect.PushMsgParams{
			Encrypted:     true,
			Encryption:    0,
			Decompression: true,
			Msg:           msg,
			ServerId:      "",
		})
		if err != nil || rsp.Code != common2.Success {
			log.Print("Push Message To Queue Err:", err)
		}
	}
}
