package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("admin")
	{
		RoleRouter.GET("/role/list")
		RoleRouter.POST("/role/add")
		RoleRouter.PUT("/role/edit")
		RoleRouter.DELETE("/role/del")
	}
}
