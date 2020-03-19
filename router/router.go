package router

import (
	"gfim/app/api/hello"
	"gfim/app/api/user"

	"gfim/app/http/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
	s.BindHandler("/login", user.SignIn)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.Group("/user/profile", func(group *ghttp.RouterGroup) {
			group.ALL("/", user.Profile)
		})
	})
}
