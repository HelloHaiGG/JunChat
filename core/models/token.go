package models

type TokenEntity struct {
	Info      *UserInfo `json:"info"`
	TimeStamp int64     `json:"time_stamp"`
}
