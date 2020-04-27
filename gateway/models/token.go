package models


type TokenEntity struct {
	Info      *UserInfo `json:"info"`
	ServerId  string    `json:"server_id"`
	TimeStamp int64     `json:"time_stamp"`
}

type UserInfo struct {
	CreateAt int64  `json:"create_at"`
	Uid      string `json:"uid" gorm:"PRIMARY_KEY;column:uid"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password" gorm:"password"`
	IsLogin  bool   `json:"is_login" gorm:"is_login"`
}
