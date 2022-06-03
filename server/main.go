package main

import (
	"admin/common"
	_ "admin/docs"
	"admin/initialize"
	_ "github.com/codyguo/godaemon" //用于后台启动时使用./admin -d true，后台持久运行
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go vben admin API
// @version 1.0
// @description  Golang admin
// @termsOfService https://github.com/admin/admin

// @contact.name liluhao
// @contact.url http://foolartist.top
// @contact.email 2415306912@qq.com

//@host 127.0.0.1:80
func main() {

	defer common.CACHE.Close()
	defer common.DB.Close()

	// 初始化路由
	routers := initialize.InitRouters()
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = routers.Run(":" + common.CONFIG.System.Port)
}

func init() {
	//初始华配置
	initialize.InitConf()
	// 初始化日志
	initialize.InitLog()
	//初始化redis
	initialize.InitCache()
	//初始化数据库
	initialize.InitDb()

	// 初始化Casbin
	initialize.InitCasbin()
}
