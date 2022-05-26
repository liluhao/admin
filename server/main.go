package server

import "admin/initialize"

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
func main() {

	//初始话化路由
	routers := initialize.InitRouters()
	routers.Run()

}
