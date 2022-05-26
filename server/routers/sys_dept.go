package routers

import "github.com/gin-gonic/gin"

func InitDeptRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("admin")
	{
		MenuRouter.GET("/dept/list")
		MenuRouter.DELETE("/dept/del")
		MenuRouter.POST("/dept/add")
		MenuRouter.PUT("/dept/edit")
	}
}
