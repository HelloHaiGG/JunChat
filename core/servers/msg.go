package servers

import (
	core "JunChat/core/protocols"
	"context"
)

type SendMessageController struct{}

func (p *SendMessageController) SendMessage(context.Context, *core.SendMsgParams) (*core.SendMsgRsp, error) {
	return nil, nil
}
