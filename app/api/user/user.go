package user

import (
	"github.com/gogf/gf/net/ghttp"
	"gfim/library/response"

)

type User struct {
	Username string `json:"username"` 
	Id       int    `json:"id"` 
	Status   string	`json:"status"` 
	Sign     string	`json:"sign"` 
	Avatar   string `json:"avatar"` 
}

type Friend struct {
	Groupname string 	`json:"groupname"` 
	Id int				`json:"id"` 
	List interface{}  	`json:"list"` 
}

type Group struct {
	Groupname string	`json:"groupname"` 
	Id int 				`json:"id"` 
	Avatar string		`json:"avatar"` 
}
func ImInitInfo(r *ghttp.Request)  {
	u1 := &User{
		Username:"飞机",
		Id: 1,
		Status: "online",
		Sign: "golang太难了",
		Avatar: "http://localhost:8199/resource/layui/css/modules/layer/default/icon.png",
	}
	u2 := &User{
		Username:"飞机2",
		Id: 2,
		Status: "online",
		Sign: "golang真的太难了",
		Avatar: "http://localhost:8199/resource/layui/css/modules/layer/default/icon.png",
	}
	friend := &Friend{
		Groupname: "我的好友",
		Id:1,
		List:[2]*User{
			u2,u2,
		},
	}
	group := &Group{
		Groupname:"Golang学习群",
		Id:1,
		Avatar:"http://localhost:8199/resource/layui/css/modules/layer/default/icon.png",
	}
	data := map[string]interface{}{
		"mine":u1,
		"friend":[2]*Friend{
			friend,friend,
		},
		"group":[4]*Group{
			group,group,group,group,
		},
	}
	response.JsonExit(r, 0, "ok", data)
}