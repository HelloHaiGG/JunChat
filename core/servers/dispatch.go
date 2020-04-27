package servers

import (
	"JunChat/common"
	"JunChat/common/iredis"
	"JunChat/config"
	"JunChat/core/models"
	"JunChat/utils"
	jsoniter "github.com/json-iterator/go"
	"sync"
)

//对节点上的链接进行计数
var DispatchMap sync.Map

func InitDispatchMap() {
	keys, _ := iredis.RedisCli.Keys("JUN:CHAT:SESSION:*").Result()
	_, _ = iredis.RedisCli.Del(keys...).Result()
	keys, _ = iredis.RedisCli.HKeys(common.LiveOnServer).Result()
	for k, _ := range config.APPConfig.JC.Nodes {
		DispatchMap.Store(k, 0)
		_, _ = iredis.RedisCli.HMSet(common.LiveOnServer, map[string]interface{}{
			k:"",
		}).Result()
	}
}

func Dispatch(uid string) (string, error) {
	serverId, err := models.GetOnlineServer(uid)
	if err != nil {
		return "", err
	}
	if serverId != "" {
		return serverId, nil
	}
	serverId = getMinLoad()
	value, _ := DispatchMap.Load(serverId)
	count := value.(int)
	DispatchMap.Store(serverId, count)
	return serverId, backUp(uid, serverId)
}

//获取最小负载
func getMinLoad() string {
	var server string
	var min int
	DispatchMap.Range(func(key, value interface{}) bool {
		if value.(int) <= min {
			min = value.(int)
			server = key.(string)
		}
		return true
	})
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
