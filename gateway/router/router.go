package router

import (
	"JunChat/common"
	common2 "JunChat/common/discover"
	"JunChat/config"
	core "JunChat/core/protocols"
	"JunChat/gateway/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/register", HandlerRegister)
	r.POST("/login", HandlerLogin)
	r.POST("/send", )

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
func HandlerSend(c *gin.Context) {

}
