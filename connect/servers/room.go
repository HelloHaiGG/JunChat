package servers

import (
	connect "JunChat/connect/protocols"
	"context"
	"sync"
)

//聊天室: 以聊天室ID为Key,用户Id数组作为值

type ChatRoomServer struct {
	Rooms sync.Map
}

func (p *ChatRoomServer) JoinRoomById(ctx context.Context, in *connect.JoinRoomParams) (*connect.JoinRoomRsp, error) {
	return nil, nil
}

func (p *ChatRoomServer) SendMsgToRoom(ctx context.Context, in *connect.SendRoomMsgParams) (*connect.SendRoomMsgRsp, error) {
	return nil, nil
}
