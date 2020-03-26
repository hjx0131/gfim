package router

import (
	"gfim/app/api/group"
	"gfim/app/api/hello"
	"gfim/app/api/user"
	"gfim/app/http/middleware"
	"gfim/app/ws/chat"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
	s.BindHandler("/signIn", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/signIn.html",
		})
	})
	s.BindHandler("/profile", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/init.html",
		})
	})
	ctlUser := new(user.Controller)
	ctlGroup := new(group.Controller)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/user/signIn", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.SignIn)
		})
		group.Middleware(middleware.Auth)
		group.Group("/user/profile", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.Profile)
		})
		group.Group("/group/userList", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlGroup.UserList)
		})
	})

	s.Group("/chat", func(group *ghttp.RouterGroup) {
		 group.Middleware(middleware.CORS)
		// group.Middleware(middleware.Auth)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", chat.WebSocket)
		})
	})
}
