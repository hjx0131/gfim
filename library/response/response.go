package response

import (
	"github.com/gogf/gf/net/ghttp"
)

//JSONResponse JSON response
type JSONResponse struct {
	Code int         `json:"code"` //错误码(0:成功,1失败)
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //返回数据(业务接口定义具体数据结构 )
}

//JSON wirte JSON
func JSON(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JSONResponse{
		Code: code,
		Msg:  msg,
		Data: responseData,
	})
}

//JSONExit wirte JSON data and exit
func JSONExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	JSON(r, err, msg, data...)
	r.Exit()
}
