package common

import "net/http"

const (
	Success = 20000

	InternalErr = http.StatusInternalServerError*100 + iota

	MissUserInfo = http.StatusForbidden*100
	UserDoesNotExistOrPasswordErr
	UserDoesNotExist
	UserAlreadyExists
	UserAlreadyLogin
	ParamsErr
)
