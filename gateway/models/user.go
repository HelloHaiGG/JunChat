package models

//用户注册
type UserRegisterParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//用户登录
type UserLoginParams struct {
	UserName string `json:"user_name"`
	Uid      string `json:"uid"`
	Password string `json:"password"`
}
