package routers

import (
	v1 "admin/app/admin"
	"admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	// 用户登录
	Router.POST("/login", v1.Login)
	Router.POST("/logout", middleware.JWTAuth(), v1.LogOut)
	Router.GET("/captcha", v1.Captcha)
	//网站搜索http://localhost/api/captcha/4FMO7B3F5St5kUsJj3Cm.png 会出现一张验证码图片
	Router.GET("/captcha/:captchaImg", v1.CaptchaImg)
}
