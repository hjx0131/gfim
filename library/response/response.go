package response

import (
	"github.com/gogf/gf/net/ghttp"
)
//数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int 		`json:"code"`  //错误码(0:成功,1失败)
	Msg 	string 		`json:"msg"`   //提示信息
	Data    interface{} `json:"data"`  //返回数据(业务接口定义具体数据结构 )
}
//标准返回结果数据结构封装
func Json(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Msg:	 msg,
		Data:    responseData,
	})
}
//返回结果并且终止程序
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}){
	Json(r, err, msg, data...)
	r.Exit()
}