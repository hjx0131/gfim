package api

import (
	"gfim/app/service/user_token"
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Base 基础控制器
type Base struct {
}

//GetIDRequest 获取登录用户ID的请求参数
type GetIDRequest struct {
	user_token.GetIDInput
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
func (b *Base) GetUserID(r *ghttp.Request) (uint, error) {
	var data *GetIDRequest
	if err := r.GetStruct(&data); err != nil {
		return 0, err
	}
	//验证token是否有效
	UserID, e := user_token.GetUserID(&data.GetIDInput)
	if e != nil {
		return 0, e
	}
	return UserID, nil
}
