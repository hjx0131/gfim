package middleware

import (
	"gfim/app/service/user_token"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//GetIDRequest 获取登录用户ID的请求参数
type GetIDRequest struct {
	user_token.GetIDInput
}

//Auth 鉴权中间件
func Auth(r *ghttp.Request) {
	var data *GetIDRequest
	if err := r.GetStruct(&data); err != nil {
		response.JSONExit(r, 2, err.Error())
	}
	//验证token是否有效
	_, err := user_token.GetUserID(&data.GetIDInput)
	if err != nil {
		response.JSONExit(r, 2, err.Error())

	}
	r.Middleware.Next()
}
