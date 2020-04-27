package common

import "net/http"

//响应相关
const (
	Success = 20000

	VerifyErr = http.StatusForbidden*100 + iota

	InternalErr = http.StatusInternalServerError*100

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
)

//程序相关
const (
	RoomChat   = 1
	SingleChat = 2
)
