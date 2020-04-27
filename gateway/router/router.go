package router

import (
	"JunChat/common"
	common2 "JunChat/common/discover"
	"JunChat/config"
	core "JunChat/core/protocols"
	"JunChat/gateway/models"
	"JunChat/utils"
	"context"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"time"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/register", HandlerRegister)
	r.POST("/login", HandlerLogin)
	r.POST("/send", SendMsg)

	return r
}

//用户注册
func HandlerRegister(c *gin.Context) {
	var params models.UserRegisterParams
	var err error
	err = c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.ParamsErr,
		})
		return
	}
	conn := common2.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewUserControllerClient(conn)
	rsp, err := client.RegisterUser(context.Background(), &core.RegisterParams{UName: params.UserName, Password: params.Password})
	if err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"Code":   common.InternalErr,
			"ErrMsg": err.Error(),
		})
		return
	}
	if rsp.Code != common.Success {
		c.JSON(http.StatusOK, &gin.H{
			"Code": rsp.Code,
		})
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"Code": common.Success,
		"Data": rsp,
	})
}

//用户登录
func HandlerLogin(c *gin.Context) {
	var params models.UserLoginParams
	var err error
	err = c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.ParamsErr,
		})
		return
	}
	conn := common2.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewUserControllerClient(conn)
	rsp, err := client.UserLogin(context.Background(), &core.LoginParams{UName: params.UserName, Password: params.Password, Uid: params.Uid})
	if err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"Code":   common.InternalErr,
			"ErrMsg": err.Error(),
		})
		return
	}
	if rsp.Code != common.Success {
		c.JSON(http.StatusOK, &gin.H{
			"Code": rsp.Code,
		})
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"Code": common.Success,
		"Data": rsp,
	})
}

//用户发送消息
func SendMsg(c *gin.Context) {
	var params models.SendMegParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.ParamsErr,
		})
		return
	}
	if !params.Check(c) {
		return
	}
	//获取token
	token := c.GetHeader("Jun-Token")
	entity := &models.TokenEntity{}
	str, err := utils.AesDecrypt(token, utils.KEY)
	if token == "" || err != nil {
		c.JSON(http.StatusForbidden, &gin.H{
			"Code": common.VerifyErr,
		})
		return
	}
	_ = jsoniter.UnmarshalFromString(str, entity)
	if time.Now().Unix()-entity.TimeStamp >= 30 {
		c.JSON(http.StatusForbidden, &gin.H{
			"Code": common.LoginTimeOut,
		})
		return
	}
	//将消息推送到Core
	conn := common2.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewSendMsgControllerClient(conn)
	rsp, err := client.SendMessage(context.Background(), &core.SendMsgParams{UserId: entity.Info.Uid, Msg: &core.MessageBody{
		Id:       params.Id,
		Text:     params.Msg,
		Urls:     params.Urls,
		SendTime: params.SendTime,
		Sender:   params.Sender,
		Receiver: entity.Info.Uid,
		MsgType:  params.Category,
	}})
	if err != nil || rsp.Code != common.Success {
		c.JSON(http.StatusOK, &gin.H{
			"Code": common.SendMsgFailed,
			"Msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, &gin.H{})
}
