package middleware

import (
	"gfim/app/service/user_token"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Auth 鉴权中间件
func Auth(r *ghttp.Request) {
	token := r.Get("token")
	if token == nil {
		response.JSONExit(r, 2, "token不能为空，请先登录")
	}
	t := token.(string)
	//验证token是否有效
	_, err := user_token.GetUserID(t)
	if err != nil {
		response.JSONExit(r, 2, err.Error())

	}
	r.Middleware.Next()
}
