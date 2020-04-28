package servers

import (
	common2 "JunChat/common"
	"JunChat/common/iredis"
	"JunChat/core/models"
	core "JunChat/core/protocols"
	"JunChat/utils"
	"context"
	jsoniter "github.com/json-iterator/go"
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

func (p *CenterServerController) OnDisconnectReport(ctx context.Context, in *core.ReportDisconnectParams) (*core.ReportDisconnectRsp, error) {

	//去除redis中的记录
	str, err := iredis.RedisCli.HGet(common2.LiveOnServer, in.ServerId).Result()
	if err != nil {
		return &core.ReportDisconnectRsp{Code: common2.ServeNotLive}, nil
	}
	users := &models.Users{}
	_ = jsoniter.UnmarshalFromString(str, users)
	i, at := utils.IncludeItem(users.Ids, in.Id)
	if !at {
		return &core.ReportDisconnectRsp{Code: common2.UserAlreadyRemove}, nil
	}
	users.Ids = append(users.Ids[:i], users.Ids[i+1:]...)
	str, _ = jsoniter.MarshalToString(users)
	suc, _ := iredis.RedisCli.HSet(common2.LiveOnServer, in.ServerId, str).Result()
	if !suc {
		return &core.ReportDisconnectRsp{Code: common2.RemoveUserIdFailed}, nil
	}
	return &core.ReportDisconnectRsp{Code: common2.Success, Id: in.Id}, nil
}
func (p *CenterServerController) OnServerChange(cxt context.Context, in *core.ReportServerStatusParams) (*core.ReportServerStatusRsp, error) {
	if in.ServerId == "" {
		return &core.ReportServerStatusRsp{Code: common2.ParamsErr}, nil
	}
	exist, _ := iredis.RedisCli.HExists(common2.LiveOnServer, in.ServerId).Result()
	if (!exist && in.Status == common2.NodeStop)||(exist && in.Status == common2.NodeStart) {
		return &core.ReportServerStatusRsp{Code: common2.Success}, nil
	}
	if in.Status == common2.NodeStart && !exist {
		ok, err := iredis.RedisCli.HSet(common2.LiveOnServer, in.ServerId, "").Result()
		if err != nil || !ok {
			return &core.ReportServerStatusRsp{Code: common2.AddNodeFailed}, nil
		}
		return &core.ReportServerStatusRsp{Code: common2.Success}, nil
	} else if in.Status == common2.NodeStop && exist {
		_, err := iredis.RedisCli.HDel(common2.LiveOnServer, in.ServerId).Result()
		if err != nil {
			return &core.ReportServerStatusRsp{Code: common2.AddNodeFailed}, nil
		}
		return &core.ReportServerStatusRsp{Code: common2.Success}, nil
	}
	return &core.ReportServerStatusRsp{Code: common2.AddNodeFailed}, nil
}
