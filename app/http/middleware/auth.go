package middleware

import (
	"gfim/app/service/user_token"
	"gfim/library/auth"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Auth 鉴权中间件
func Auth(r *ghttp.Request) {
	token := auth.GetToken(r)
	//验证token是否有效
	_, e := user_token.GetUserID(&user_token.GetIDInput{
		Token: token,
	})
	if e != nil {
		response.JSONExit(r, 2, e.Error())
	}
	r.Middleware.Next()
}
