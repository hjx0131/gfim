package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//SetData 模版注入公共数据
func SetData(r *ghttp.Request) {
	url := g.Cfg().Get("websocket.Address")
	if url != nil {
		url = url.(string)
	}
	r.Response.WriteTpl("header.html", g.Map{
		"wsURL": url,
	})
	r.Middleware.Next()
}
