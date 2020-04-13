package user

import (
	"gfim/app/api"
	"gfim/app/service/user"
	"gfim/app/service/user_token"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

// SignInRequest 登录请求参数
type SignInRequest struct {
	Username string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

//SignUpRequest 注册请求参数
type SignUpRequest struct {
	user.SignUpInput
}

//SignIn 登录
func (c *Controller) SignIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	token, err := user.SignIn(data.Username, data.Password)
	if err != nil {
		c.Fail(r, err.Error())
	}
	resp := map[string]string{
		"token": token,
	}
	c.Success(r, resp)
}

//SignUp 注册
func (c *Controller) SignUp(r *ghttp.Request) {
	var data *SignUpRequest
	if err := r.GetStruct(&data); err != nil {
		c.Fail(r, err.Error())
	}
	err := user.SignUp(&data.SignUpInput, r.GetClientIp())
	if err != nil {
		c.Fail(r, err.Error())
	}
	c.Success(r, nil)
}

//Logout 注销登录
func (c *Controller) Logout(r *ghttp.Request) {
	type logoutRequest struct {
		user_token.GetIDInput
	}
	var data *logoutRequest
	if err := r.GetStruct(&data); err != nil {
		c.Fail(r, err.Error())
	}
	user_token.Logout(&data.GetIDInput)
	c.Success(r, nil)
}

//Profile 主面板
func (c *Controller) Profile(r *ghttp.Request) {
	id := c.GetUserID(r)
	data, e := user.Profile(id)
	if e != nil {
		c.Fail(r, e.Error())
	}
	c.Success(r, data)
}

//Search 搜索用户
func (c *Controller) Search(r *ghttp.Request) {
	id := c.GetUserID(r)
	var data *user.SearchRequst
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	resp, err := user.Search(data, id)
	if err != nil {
		c.Fail(r, err.Error())
	}
	c.Success(r, resp)
}

//Recommend 推荐用户
func (c *Controller) Recommend(r *ghttp.Request) {
	id := c.GetUserID(r)
	var data *user.SearchRequst
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	resp, err := user.Recommend(data, id)
	if err != nil {
		c.Fail(r, err.Error())
	}
	c.Success(r, resp)
}
