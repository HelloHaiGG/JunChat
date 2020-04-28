package servers

import (
	"JunChat/common"
	connect "JunChat/connect/protocols"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
)

type PushMessageController struct {
}

func (p *PushMessageController) PushMsgToConnectServer(ctx context.Context, in *connect.PushMsgParams) (*connect.PushMsgRsp, error) {

	//TODO 判断是否解密
	//TODO 判断是否解压缩
	//判断是否推送到了正确的服务
	//if in.ServerId != NETServer {
	//	log.Error("服务节点错误.")
	//	return &connect.PushMsgRsp{Code: common.SendMsgFailed}, nil
	//}
	//获取用户链接
	if in.Msg.Receiver == "" {
		return &connect.PushMsgRsp{Code: common.SendMsgFailed}, nil
	}
	conn, ok := HandleConn.Load(in.Msg.Receiver)
	if !ok {
		log.Error("用户未链接.")
		return &connect.PushMsgRsp{Code: common.SendMsgFailed}, nil
	}
	c, _ := conn.(*Connect)
	b, _ := json.Marshal(in.Msg)
	_ = c.Conn.WriteMessage(websocket.TextMessage, b)

	return &connect.PushMsgRsp{Code:common.Success}, nil
}
