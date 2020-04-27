package models

import (
	"JunChat/common"
	"JunChat/config"
	core "JunChat/core/protocols"
	"errors"
	jsoniter "github.com/json-iterator/go"
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
	}
	if p.MsgType == common.SingleChat {
		server, err = GetOnlineServer(p.Receiver)
		if err != nil {
			return "", errors.New("Not Find Server Id. ")
		}
		if server == "" {
			//TODO 离线消息
			return "",errors.New("User Not OnLine. ")
		}
		p.N[p.Receiver] = &Node{
			Host: "127.0.0.1",
			Port: config.APPConfig.JC.Nodes[server],
		}
	}
	str, _ := jsoniter.MarshalToString(p)
	return str, nil
}
