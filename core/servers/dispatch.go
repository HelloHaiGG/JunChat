package servers

import (
	"JunChat/common"
	common2 "JunChat/common/discover"
	"JunChat/common/iredis"
	"JunChat/config"
	"JunChat/core/models"
	"JunChat/utils"
	"context"
	jsoniter "github.com/json-iterator/go"
)

func InitDispatchMap() {
	keys, _ := iredis.RedisCli.Keys("JUN:CHAT:SESSION:*").Result()
	_, _ = iredis.RedisCli.Del(keys...).Result()
	keys, _ = iredis.RedisCli.HKeys(common.LiveOnServer).Result()
	for k, v := range config.APPConfig.JC.Nodes {
		if !DialConnectServer(v) {
			continue
		} else {
			ok, _ := iredis.RedisCli.HExists(common.LiveOnServer, k).Result()
			if ok {
				continue
			}
		}
		_, _ = iredis.RedisCli.HMSet(common.LiveOnServer, map[string]interface{}{
			k: "",
		}).Result()
	}
}

func Dispatch(uid string) (string, error) {
	serverId, err := models.GetOnlineServer(uid)
	if err != nil || serverId != "" {
		return serverId, err
	}
	serverId = getMinLoad()
	return serverId, backUp(uid, serverId)
}

//获取最小负载
func getMinLoad() string {
	var server string
	var min int
	result, _ := iredis.RedisCli.HGetAll(common.LiveOnServer).Result()
	for key, value := range result {
		if value == "" {
			return key
		} else {
			users := models.Users{}
			_ = jsoniter.UnmarshalFromString(value, users)
			if len(users.Ids) <= min {
				min = len(users.Ids)
				server = key
			}
		}
	}
	return server
}

//将用户所在server备份到redis
func backUp(uid, server string) error {
	count, err := iredis.RedisCli.HLen(common.LiveOnServer).Result()
	if err != nil {
		return err
	}
	users := &models.Users{}
	if count == 0 {
		users.Ids = []string{uid}
	} else {
		str, _ := iredis.RedisCli.HGet(common.LiveOnServer, server).Result()
		_ = jsoniter.UnmarshalFromString(str, users)
		if _, in := utils.IncludeItem(users.Ids, uid); !in {
			users.Ids = append(users.Ids, uid)
		}
	}
	str, _ := jsoniter.MarshalToString(users)
	_, err = iredis.RedisCli.HSet(common.LiveOnServer, server, str).Result()
	return err
}

//先检查服务是否启动
func DialConnectServer(port string) bool {
	conn := common2.GetServerConnByHost("127.0.0.1", port)
	client := common.NewProtoDialClient(conn)

	//3s
	//cxt, _ := context.WithTimeout(context.Background(), time.Second*3)

	rsp, err := client.TryDial(context.Background(), &common.Request{
		Ping: "Core Server",
	})
	if err != nil || rsp.Code != common.Success {
		return false
	}
	return true
}
