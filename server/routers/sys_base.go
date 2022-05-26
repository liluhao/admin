package routers

import "github.com/gin-gonic/gin"

func InitBaseRouter(Router *gin.RouterGroup) {
	// 用户登录
	Router.POST("/login")
	Router.POST("/logout")

	Router.GET("/captcha")
	Router.GET("/captcha/:captchaImg")

}
