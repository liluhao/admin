package initialize

import (
	"admin/routers"
	"github.com/gin-gonic/gin"
)

func InitRouters() (Router *gin.Engine) {
	Router = gin.Default()

	ApiV1Group := Router.Group("/api")
	routers.InitBaseRouter(ApiV1Group)
	routers.InitUserRouter(ApiV1Group)
	routers.InitRoleRouter(ApiV1Group)
	routers.InitMenuRouter(ApiV1Group)
	routers.InitDeptRouter(ApiV1Group)
	return
}
