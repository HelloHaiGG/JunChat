package models

import (
	"JunChat/common"
	"JunChat/config"
	core "JunChat/core/protocols"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type MsgWrap struct {
	*core.MessageBody
	N map[string]*Node `json:"N"` //接受者所在节点信息
}

type Node struct {
	Host string `json:"Host,omitempty"`
	Port string `json:"Port,omitempty"`
}

func (p *MsgWrap) WrapNode(in *core.SendMsgParams) (string, error) {
	var err error
	var server string
	p.MessageBody = in.Msg
	if p.MsgType == common.RoomChat {
		//聊天室消息的 receiver id 为聊天室的id
		//通过聊天室ID,获取聊天室成员,然后再奉封装成员所在的服务器信息
	}
	if p.MsgType == common.SingleChat {
		server, err = GetOnlineServer(p.Receiver.Uid)
		if err != nil {
			return "", errors.New("Not Find Server Id. ")
		}
		if server == "" {
			//TODO 离线消息
			return "", errors.New("User Not OnLine. ")
		}
		p.N[p.Receiver.Uid] = &Node{
			Host: "127.0.0.1",
			Port: config.APPConfig.CN.Nodes[server],
		}
	}

	//封装发送者用户基本信息
	info, err := GetUserInfoById(p.Sender.Uid)
	if err != nil || info == nil {
		log.Printf("Get Sender Info Err: %v", err)
	} else {
		p.Sender.UserName = info.UserName
	}

	str, _ := jsoniter.MarshalToString(p)
	return str, nil
}
