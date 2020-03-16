package router

import (
	"gfim/app/api/hello"
	"gfim/app/api/user"
    "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
	s.BindHandler("/imInitInfo", user.ImInitInfo)

}
