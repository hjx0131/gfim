package user

import (
	"gfim/app/service/user"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// SignInRequest 登录请求参数
type SignInRequest struct {
	Username string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

//SignIn 登录
func SignIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 1, err.Error())

	}
	token, err := user.SignIn(data.Username, data.Password)
	if err != nil {
		response.JSONExit(r, 1, err.Error())

	}
	resp := map[string]string{
		"token": token,
	}
	response.JSONExit(r, 0, "ok", resp)

}

//Profile 主面板
func Profile(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", "主面板")

}
