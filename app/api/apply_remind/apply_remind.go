package apply_remind

import (
	"gfim/app/api"
	"gfim/app/service/apply_remind"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

//SetIsRead 根据用户ID将提醒标记为已读
func (c *Controller) SetIsRead(r *ghttp.Request) {
	apply_remind.SetIsRead(c.GetUserID(r))
	c.Success(r)
}
