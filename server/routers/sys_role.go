package routers

import (
	v1 "admin/app/admin"
	"admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("admin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.Logger())
	{
		RoleRouter.GET("/role/list", v1.GetRoleList) //guest测试账号可以获取角色列表
		RoleRouter.POST("/role/add", v1.CreateRole)
		RoleRouter.PUT("/role/edit", v1.UpdateRole)
		RoleRouter.DELETE("/role/del", v1.DelRole)
	}
}
