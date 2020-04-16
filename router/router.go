package router

import (
	"gfim/app/api/group"
	"gfim/app/api/record"
	"gfim/app/api/user"

	"gfim/app/http/middleware"
	"gfim/app/ws/chat"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/signIn", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/signIn.html",
		})
	})
	s.BindHandler("/signUp", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/signup.html",
		})
	})
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/index.html",
		})
	})
	s.BindHandler("/chatlog", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/chatlog.html",
		})
	})
	s.BindHandler("/msgbox", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/msgbox.html",
		})
	})
	s.BindHandler("/find", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"mainTpl": "index/find.html",
		})
	})
	ctlUser := new(user.Controller)
	ctlGroup := new(group.Controller)
	ctlRecord := new(record.Controller)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/user/signIn", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.SignIn)
		})
		group.Group("/user/signUp", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.SignUp)
		})
		group.Group("/user/logout", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.Logout)
		})
		group.Middleware(middleware.Auth)
		group.Group("/user/profile", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.Profile)
		})
		group.Group("/user/recommend", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.Recommend)
		})
		group.Group("/user/search", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlUser.Search)
		})
		group.Group("/group/userList", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlGroup.UserList)
		})
		group.Group("/record/getData", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlRecord.GetData)
		})
	})
	ctlChat := new(chat.Controller)

	s.Group("/chat", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		// group.Middleware(middleware.Auth)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlChat.WebSocket)
		})
	})
}
