package models

type TokenEntity struct {
	Info      *UserInfo `json:"info"`
	ServerId  string    `json:"server_id"`
	TimeStamp int64     `json:"time_stamp"`
}

type HeartBeat struct {
	Msg string `json:"msg"`
}