package servers

import (
	"JunChat/common"
	"JunChat/utils"
	"context"
)

type GateWayDialServer struct{}

func (p *GateWayDialServer) TryDial(ctx context.Context, in *common.Request) (*common.Response, error) {
	ip := utils.GetInternalIp()
	return &common.Response{Pong: "GateWay Running:" + ip}, nil
}
