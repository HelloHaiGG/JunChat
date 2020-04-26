package servers

import (
	connect "JunChat/connect/protocols"
	"context"
	"google.golang.org/grpc"
)

type PushMessageController struct {
}

func (p *PushMessageController) PushMsgToConnectServer(ctx context.Context, in *connect.PushMsgParams, opts ...grpc.CallOption) (*connect.PushMsgRsp, error) {
	return nil, nil
}
