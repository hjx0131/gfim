package apply

import (
	"gfim/app/api"
	"gfim/app/service/apply"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

//GetDataReq GetData请求参数格式
type GetDataReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

//GetData 获取列表和总数
func (c *Controller) GetData(r *ghttp.Request) {
	var data *GetDataReq
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	res, err := apply.GetListAndTotal(&apply.GetListRequest{
		UserID: c.GetUserID(r),
		Page:   data.Page,
		Limit:  data.Limit,
	})
	if err != nil {
		c.Fail(r, err.Error())
	}
	c.Success(r, res)
}
