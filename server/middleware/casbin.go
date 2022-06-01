package middleware

import (
	"admin/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetClaims(c)
		//获取用户的角色
		sub := user.RoleId
		//获取请求的URI
		obj := strings.Replace(c.Request.URL.Path, "/api", "", 1)
		//fmt.Printf("%q\n", c.Request.URL.Path) //打印:"/api/admin/user/add"
		//fmt.Printf("%q\n", obj) //打印:"/admin/user/add"

		//获取请求方法
		act := c.Request.Method

		_ = common.CASBIN.LoadPolicy()
		ok, err := common.CASBIN.EnforceSafe(sub, obj, act)
		//fmt.Println(ok)  //false
		//fmt.Println(err) //<nil>
		//product模式需要经过authentication，develop不需要经过authentication

		if common.CONFIG.System.Env == "develop" || ok {
			c.Next()
		} else {
			if err != nil {
				common.LOG.Warn("Casbin规则错误" + "\t" + err.Error())
			}
			//[admin] 2022/05/31 - 22:31:03   warn    guest   /admin/user/add POST
			common.LOG.Warn(user.Username + "\t" + obj + "\t" + act)
			//前端返回时不会解析\t,仍然返回：{"code":403,"msg":"非法请求\t权限不足","result":null}
			ResponseFail(c, 403, "非法请求\t权限不足")
			//fmt.Println("非法请求\t权限不足")打印：非法请求        权限不足
			c.Abort()
		}
	}
}
