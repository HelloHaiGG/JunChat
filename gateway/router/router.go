package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/register", HandlerRegister)
	r.GET("/chat", HandlerLogin)

	return r
}

//用户注册
func HandlerRegister(c *gin.Context) {

}

//用户登录
func HandlerLogin(c *gin.Context) {

}
