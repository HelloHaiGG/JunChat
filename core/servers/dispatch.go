package servers

import (
	"JunChat/common/iredis"
	"JunChat/config"
	"JunChat/utils"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"sync"
)

//对节点上的链接进行计数
var dispatchMap sync.Map

func InitDispatchMap() {
	for i, _ := range config.APPConfig.CN.Nodes {
		dispatchMap.Store(fmt.Sprintf("core-node-%d", i+1), 0)
	}
}

func Dispatch(uid string) (string, error) {
	serverId ,err := GetOnlineServer(uid)
	if err != nil {
		return "",err
	}
	if serverId != ""{
		return serverId,nil
	}
	serverId = getMinLoad()
	value, _ := dispatchMap.Load(serverId)
	count := value.(int)
	dispatchMap.Store(serverId, count)
	return serverId, backUp(uid, serverId)
}

//获取最小负载
func getMinLoad() string {
	var server string
	var min int
	dispatchMap.Range(func(key, value interface{}) bool {
		if value.(int) <= min {
			min = value.(int)
			server = key.(string)
		}
		return true
	})
	return server
}

type Users struct {
	Ids []string `json:"ids"`
}

//将用户所在server备份到redis
func backUp(uid, server string) error {
	count, err := iredis.RedisCli.HLen("SERVER:USER").Result()
	if err != nil {
		return err
	}
	users := &Users{}
	if count == 0 {
		users.Ids = []string{uid}
	} else {
		str, _ := iredis.RedisCli.HGet("SERVER:USER", server).Result()
		_ = jsoniter.UnmarshalFromString(str, users)
		if _, in := utils.IncludeItem(users.Ids, uid); !in {
			users.Ids = append(users.Ids, uid)
		}
	}
	str, _ := jsoniter.MarshalToString(users)
	_, err = iredis.RedisCli.HSet("SERVER:USER", server, str).Result()
	return err
}

//遍历所有server,如果现在直接返回所在serverId
func GetOnlineServer(uid string) (string, error) {
	server, err := iredis.RedisCli.HGetAll("SERVER:USER").Result()
	if err != nil {
		return "", err
	}
	for str, _ := range server {
		users := &Users{}
		_ = jsoniter.UnmarshalFromString(server[str], users)
		if _, in := utils.IncludeItem(users.Ids, uid); in {
			return str, nil
		}
	}
	return "", nil
}
