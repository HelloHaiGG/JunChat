package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5WithStr(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	str = hex.EncodeToString(b)
	return str
}
