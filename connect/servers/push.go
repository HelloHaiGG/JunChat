package servers

import (
	"JunChat/common"
	connect "JunChat/connect/protocols"
	"context"
)

type PushMessageController struct {
}

func (p *PushMessageController) PushMsgToConnectServer(ctx context.Context, in *connect.PushMsgParams) (*connect.PushMsgRsp, error) {

	//TODO 判断是否解密
	//TODO 判断是否解压缩
	//TODO 判断是否推送到了正确的服务

	//TODO 获取用户链接
	if in.Msg.Receiver == "" {
		return &connect.PushMsgRsp{Code: common.SendMsgFailed}, nil
	}
	conn, ok := HandleConn.Load(in.Msg.Receiver)
	if !ok {

	}
	c, _ := conn.(*Connect)
	_ = c.Conn.WriteMessage(0, nil)

	return nil, nil
}
