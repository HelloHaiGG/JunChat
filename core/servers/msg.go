package servers

import (
	"JunChat/common"
	"JunChat/common/iredis"
	"JunChat/core/models"
	core "JunChat/core/protocols"
	"context"
	"github.com/prometheus/common/log"
)

type SendMessageController struct{}

func (p *SendMessageController) SendMessage(cxt context.Context, in *core.SendMsgParams) (*core.SendMsgRsp, error) {
	msg :=new(models.MsgWrap)
	msg.N = map[string]*models.Node{}
	body, err := msg.WrapNode(in)
	if err != nil {
		log.Error("Wrap Node Err:", err)
		return &core.SendMsgRsp{Code: common.SendMsgFailed}, nil
	}
	_, err = iredis.RedisCli.LPush(common.MsgQueue, body).Result()
	if err != nil {
		log.Error("Push Msg To Queue Err:", err)
		return &core.SendMsgRsp{Code: common.SendMsgFailed}, nil
	}
	log.Info("Push Msg To Queue Suc. Msg Id:",msg.Id)
	return &core.SendMsgRsp{Code: common.Success}, nil
}
