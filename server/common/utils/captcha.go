package utils

import (
	"admin/common"
	"bytes"
	"github.com/dchest/captcha"
	"net/http"
	"path"
	"strings"
	"time"
)

func CaptchaServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL)      //   api/captcha/8MbU9p2hZUCYoEplc8qD.png
	//fmt.Println(r.URL.Path) //   api/captcha/8MbU9p2hZUCYoEplc8qD.png
	//Split函数将路径从最后一个斜杠后面位置分隔为两个部分并返回，path == dir+file
	dir, file := path.Split(r.URL.Path)
	//fmt.Println(dir) //   /api/captcha/
	//fmt.Println(file)//   8MbU9p2hZUCYoEplc8qD.png
	/*
			file := "aaa.bbb.vvvn"
		    ext := path.Ext(file)   //.vvvn
		    id := file[:len(file)-len(ext)]//aaa.bbb
	*/
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	//fmt.Println(ext) //   .png
	//fmt.Println(id)  //    8MbU9p2hZUCYoEplc8qD
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	//fmt.Println(r.FormValue("reload"))//  空行
	//fmt.Println(r.FormValue("lang"))//  空行
	//fmt.Println(path.Base(dir)) //    captcha
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	//fmt.Println(download) //false
	if Serve(w, r, id, ext, lang, download, common.CONFIG.Captcha.ImgWidth, common.CONFIG.Captcha.ImgHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

//强制浏览器不使用缓存
//NoCache 是一个中间件函数，它附加标头以防止客户端缓存 HTTP 响应。
//NOCache是请求头与响应头里的字段；通过指定指令来实现缓存机制。缓存指令是单向的，这意味着在请求中设置的指令，不一定被包含在响应中
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	//no-cache:在发布缓存副本之前，强制要求缓存把请求提交给原始服务器进行验证(协商缓存验证)。
	//no-store ：缓存不应存储有关客户端请求或服务器响应的任何内容，即不使用任何缓存。
	//must-revalidate:一旦资源过期（比如已经超过），在成功向原始服务器验证之前(即你也可以把 revalidate 理解成“再次校验”的意思：再次校验看看缓存是不是真的过期了)，缓存不能用该资源响应后续请求
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//与 Cache-Control: no-cache 效果一致。强制要求缓存服务器在返回缓存的版本之前将请求提交到源头服务器进行验证。
	w.Header().Set("Pragma", "no-cache")
	//一个 HTTP-日期 时间戳,即在此时候之后，响应过期;如果在Cache-Control响应头设置了 "max-age" 或者 "s-max-age" 指令，那么头会被忽略；0 ：代表着过去的日期，即该资源已经过期。
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	//前端的验证码图片是后端产生的，下面这行代码用于把图片传给前端
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
