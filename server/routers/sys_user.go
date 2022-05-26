package routers

import "github.com/gin-gonic/gin"

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("admin")
	{
		UserGroup.GET("/user/list")
		UserGroup.POST("/user/add")
		UserGroup.PUT("/user/edit")
		UserGroup.DELETE("/user/del")
		UserGroup.GET("/user/enu")
		UserGroup.PUT("/password")
	}
}
