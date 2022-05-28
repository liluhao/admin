package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//设置响应头信息
		//Access-Control-Allow-Origin:响应应头指定了该响应的资源是否被允许与给定的origin共享(origin其实是一个url，当协议、主机和端口都匹配时,两个对象就具有相同的origin。)，*代表允许所有域都具有访问该响应的资源权限
		c.Header("Access-Control-Allow-Origin", "*")
		//Access-Control-Allow-Headers：用于预检请求中，列出了将会在正式请求的Access-Control-Request-Headers字段中出现的首部信息。
		//Access-Control-Request-Headers: 出现于预检请求中，用于通知服务器在真正的请求中会采用哪些请求头。
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		//Access-Control-Allow-Methods:对预检请求的应答中明确了客户端所要访问的资源允许使用的方法或方法列表
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法
		//Access-Control-Expose-Headers:列出了哪些首部可以作为响应的一部分暴露给外部。
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//Access-Control-Allow-Credentials:是否可以将对请求的响应暴露给页面。返回true则可以，其他值均不可以。 Credentials可以是 cookies, authorization headers 或 TLS client certificates。
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
