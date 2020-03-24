package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

//CORS 允许接口跨域请求
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
