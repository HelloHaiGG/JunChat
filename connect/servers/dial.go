package servers

import (
	"JunChat/common"
	"JunChat/utils"
	"context"
)

type ConnectDialServer struct{}

func (p *ConnectDialServer) TryDial(ctx context.Context, in *common.Request) (*common.Response, error) {
	ip := utils.GetInternalIp()
	return &common.Response{Pong: "Core Running:" + ip,Code:common.Success}, nil
}
