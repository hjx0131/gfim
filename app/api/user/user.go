package user

import (
	"gfim/app/api"
	"gfim/app/service/user"

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

//Profile 主面板
func (c *Controller) Profile(r *ghttp.Request) {
	ID := c.GetUserID(r)
	data, e := user.Profile(ID)
	if e != nil {
		c.Fail(r, e.Error())
	}
	c.Success(r, data)
}
