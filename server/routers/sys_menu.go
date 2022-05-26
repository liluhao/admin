package routers

import "github.com/gin-gonic/gin"

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("admin")
	{
		MenuRouter.GET("/menu/list")
		MenuRouter.POST("/menu/add")
		MenuRouter.PUT("/menu/edit")
		MenuRouter.DELETE("/menu/del")
	}
}
