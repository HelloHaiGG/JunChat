package models

import (
	"JunChat/common"
	"JunChat/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

type SendMegParams struct {
	Id       string   `json:"id"`
	Category int32    `json:"category"` //消息种类
	Sender   string   `json:"sender"`
	Receiver string   `json:"receiver"`
	Msg      string   `json:"msg"`
	SendTime int64    `json:"send_time"`
	Urls     []string `json:"urls"`
}

func (p *SendMegParams) Check(c *gin.Context) bool {
	p.Id = cast.ToString(utils.SFIdTool.GetID())
	if p.Category == 0 {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.ParamsErr,
		})
		return false
	}
	if p.Receiver == "" {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.ParamsErr,
		})
		return false
	}
	p.Msg = strings.TrimSpace(p.Msg)
	p.SendTime = time.Now().Unix()
	return true
}
