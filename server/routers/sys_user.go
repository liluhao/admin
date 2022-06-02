package routers

import (
	v1 "admin/app/admin"
	"admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("admin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.Logger())
	{
		UserGroup.GET("/user/list", v1.GetUserList) //chris测试账号可以获取用户列表
		UserGroup.POST("/user/add", v1.CreateUser)  //chris测试账号不能添加账号,会给前端下响应ResponseFail(c, 403, "非法请求\t权限不足")，页面弹出”对不起，您没有权限执行该操作！“
		UserGroup.PUT("/user/edit", v1.UpdateUser)  //chris测试账号不能更新账号信息,自己的账户信息也不能改
		UserGroup.DELETE("/user/del", v1.DelUser)
		//下面api使用到缓存
		UserGroup.GET("/user/menu", v1.GetUserMenuByClaims) //该api用于再登录页面点击login后进行跳转时跳转
		UserGroup.PUT("/password", v1.ResetPwd)
	}
}
