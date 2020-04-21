package api

import (
	"gfim/app/service/user_token"
	"gfim/library/auth"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Base 基础控制器
type Base struct {
}

//Success 成功数据返回
func (b *Base) Success(r *ghttp.Request, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	response.JSONExit(r, 0, "ok", responseData)
}

//Fail 失败数据返回
func (b *Base) Fail(r *ghttp.Request, msg string) {
	response.JSONExit(r, 1, msg)
}

//GetUserID 获取登录用户ID
func (b *Base) GetUserID(r *ghttp.Request) uint {
	token := auth.GetToken(r)
	//验证token是否有效
	userID, e := user_token.GetUserID(&user_token.GetIDInput{
		Token: token,
	})
	if e != nil {
		response.JSONExit(r, 2, e.Error())
	}
	return userID
}
