package models

import (
	"JunChat/common"
	connect "JunChat/connect/protocols"
	"encoding/json"
	"time"
)

func GetDailyMessage() []byte {
	body := connect.MessageBody{
		Text:     "This JunChat!  ·^_^· ",
		SendTime: time.Now().Unix(),
		MsgType:  common.System,
	}
	b, _ := json.Marshal(body)
	return b
}
