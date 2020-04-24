package friend_group

import (
	"gfim/app/api"
	"gfim/app/service/friend_group"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

//SaveInpnut 保存请求参数
type SaveInpnut struct {
	friend_group.SaveReq
}

//Save 保存数据
func (c *Controller) Save(r *ghttp.Request) {
	var data *SaveInpnut
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	data.SaveReq.UserID = c.GetUserID(r)
	if err := friend_group.Save(&data.SaveReq); err != nil {
		c.Fail(r, err.Error())
	}
	c.Success(r)
}
