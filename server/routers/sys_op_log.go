package routers

import (
	v1 "admin/app/admin"
	"admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOpLogRouter(Router *gin.RouterGroup) {
	OpLogRouter := Router.Group("admin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OpLogRouter.GET("/log/list", v1.GetOperLogList)
		OpLogRouter.DELETE("/log/del", v1.DeleteLog)
		OpLogRouter.DELETE("/log/del_batch", v1.BatchDeleteLog)
	}
}
