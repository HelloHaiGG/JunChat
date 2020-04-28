package common

import "net/http"

//响应相关
const (
	Success = 20000

	VerifyErr = http.StatusForbidden*100 + iota

	InternalErr = http.StatusInternalServerError * 100

	MissUserInfo = http.StatusForbidden * 100
	UserDoesNotExistOrPasswordErr
	UserDoesNotExist
	UserAlreadyExists
	UserAlreadyLogin
	LoginTimeOut
	ParamsErr

	ServeNotLive
	UserAlreadyRemove
	RemoveUserIdFailed

	//消息发送失败
	SendMsgFailed = 60001
	//没有可用的server
	NoUsableServer = 60002
	//添加活跃节点失败
	AddNodeFailed = 60003
)

//程序相关
const (
	RoomChat   = 1
	SingleChat = 2

	//玩家所在节点
	LiveOnServer = "LIVE:ON:SERVER"

	//消息队列
	MsgQueue = "JUN:CHAT:MSG:CHAN"

	//节点开启
	NodeStart = 1
	NodeStop  = 2
)
