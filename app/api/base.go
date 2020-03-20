package api

import (
	"gfim/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Base 基础控制器
type Base struct {
}

//Success 成功数据返回
func (b *Base) Success(r *ghttp.Request, data ...interface{}) {
	response.JSONExit(r, 0, "ok", data)

}

//Fail 失败数据返回
func (b *Base) Fail(r *ghttp.Request, msg string) {
	response.JSONExit(r, 1, msg)
}
