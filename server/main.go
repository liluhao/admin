package server

import "admin/initialize"

func main() {

	//初始话化路由
	routers := initialize.InitRouters()
	routers.Run()
}
