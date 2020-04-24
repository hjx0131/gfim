package router

import (
	"gfim/app/api/apply"
	"gfim/app/api/apply_remind"
	"gfim/app/api/friend_group"
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

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.SetData)
		group.ALL("/signIn", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/signIn.html",
			})
		})
		group.ALL("/signUp", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/signup.html",
			})
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/index.html",
			})
		})
		group.ALL("/chatlog", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/chatlog.html",
			})
		})
		group.ALL("/msgbox", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/msgbox.html",
			})
		})
		group.ALL("/find", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/find.html",
			})
		})
		group.ALL("/createFriendGroup", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/createFriendGroup.html",
			})
		})
		group.ALL("/createGroup", func(r *ghttp.Request) {
			r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "index/createGroup.html",
			})
		})
	})

	ctlUser := new(user.Controller)
	ctlGroup := new(group.Controller)
	ctlRecord := new(record.Controller)
	ctlApply := new(apply.Controller)
	ctlApplyRemind := new(apply_remind.Controller)
	ctlFriendGroup := new(friend_group.Controller)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.ALL("/signIn", ctlUser.SignIn)
			group.ALL("/signUp", ctlUser.SignUp)
			group.Middleware(middleware.Auth)
			group.ALL("/logout", ctlUser.Logout)
			group.ALL("/profile", ctlUser.Profile)
			group.ALL("/recommend", ctlUser.Recommend)
			group.ALL("/search", ctlUser.Search)
		})
		group.Middleware(middleware.Auth)
		group.Group("/group", func(group *ghttp.RouterGroup) {
			group.ALL("/userList", ctlGroup.UserList)
			group.ALL("/save", ctlGroup.Save)
			group.ALL("/search", ctlGroup.Search)


		})
		group.Group("/record", func(group *ghttp.RouterGroup) {
			group.ALL("/getData", ctlRecord.GetData)
		})
		group.Group("/apply", func(group *ghttp.RouterGroup) {
			group.ALL("/getData", ctlApply.GetData)
		})
		group.Group("/applyRemind", func(group *ghttp.RouterGroup) {
			group.ALL("/setIsRead", ctlApplyRemind.SetIsRead)
		})
		group.Group("/friendGroup", func(group *ghttp.RouterGroup) {
			group.ALL("/save", ctlFriendGroup.Save)
		})
	})
	ctlChat := new(chat.Controller)

	s.Group("/chat", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctlChat.WebSocket)
		})
	})
}
