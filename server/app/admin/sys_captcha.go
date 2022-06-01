package admin

import (
	"admin/common"
	"admin/common/utils"
	"admin/middleware"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func Captcha(c *gin.Context) {
	//KeyLong在配置文件中固定为4
	//fmt.Println(common.CONFIG.Captcha.KeyLong)//恒打印出4
	//NewLen会随机产生id,每次产生的都不一样
	captchaId := captcha.NewLen(common.CONFIG.Captcha.KeyLong)
	//fmt.Println(captchaId) //Kdm4iAZV7woMY3oQ0g10
	//举例:  {"code":200,"msg":"验证码获取成功","result":{"CaptchaId":"t96fFKkIHmt4V53sNWBc","CaptchaSrc":"http://localhost/api/captcha/t96fFKkIHmt4V53sNWBc.png"}}
	middleware.ResponseSucc(c, "验证码获取成功", map[string]interface{}{
		"CaptchaId":  captchaId,
		"CaptchaSrc": "http://" + c.Request.Host + c.Request.URL.Path + "/" + captchaId + ".png",
	})
}

func CaptchaImg(c *gin.Context) {
	utils.CaptchaServeHTTP(c.Writer, c.Request)
}
