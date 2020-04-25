package servers

import (
	"JunChat/common"
	"JunChat/utils"
	"context"
)

type QueueDialServer struct {}

func (p *QueueDialServer) TryDial(ctx context.Context, in *common.Request) (*common.Response, error) {
	ip := utils.GetInternalIp()
	return &common.Response{Pong: "Queue Running:" + ip}, nil
}