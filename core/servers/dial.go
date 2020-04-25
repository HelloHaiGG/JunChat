package servers

import (
	"JunChat/common"
	"JunChat/utils"
	"context"
)

type CoreDialServer struct{}

func (p *CoreDialServer) TryDial(ctx context.Context, in *common.Request) (*common.Response, error) {
	ip := utils.GetInternalIp()
	return &common.Response{Pong: "Core Running:" + ip}, nil
}
