package group

import (
	"gfim/app/api"
	"gfim/app/service/group"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

//UserListRequest 群员列表请求参数
type UserListRequest struct {
	ID uint `v:"required#群组ID为空"`
}

//UserList 获取群员列表
func (c *Controller) UserList(r *ghttp.Request) {
	var data *UserListRequest
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())

	}
	list, err := group.GetUserListByID(data.ID)
	if err != nil {
		c.Fail(r, err.Error())
	}
	res := gmap.New()
	res.Sets(map[interface{}]interface{}{
		"list": list,
	})
	c.Success(r, res)
}
