package servers

import (
	core "JunChat/core/protocols"
	"context"
	"sync"
)

type Entity struct {
	Id       string
	Category int32 //1.用户,2.聊天室
}

type CenterServerController struct {
	Record sync.Map
}

func (p *CenterServerController) GetServer(ctx context.Context, in *core.GetServerByIdParams) (*core.GetServerByIdRsp, error) {
	return nil, nil
}

func (p *CenterServerController) GetRoomMembers(ctx context.Context, in *core.GetMembersParams) (*core.GetMembersRsp, error) {
	return nil, nil
}

func (p *CenterServerController) Report(ctx context.Context, in *core.ReportLogoutParams) (*core.ReportLogoutRsp, error) {
	return nil, nil
}
