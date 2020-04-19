package record

import (
	"gfim/app/api"
	"gfim/app/service/group_record"
	"gfim/app/service/user_record"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 控制器结构体
type Controller struct {
	api.Base
}

//GetRecordRequest 获取记录所需要的参数
type GetRecordRequest struct {
	Type  string `v:"in:friend,group#类型错误" json:"type"`
	ID    uint   `v:"required#ID不能为空" json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

//GetData 获取列表和总数
func (c *Controller) GetData(r *ghttp.Request) {
	var data *GetRecordRequest
	if err := r.Parse(&data); err != nil {
		c.Fail(r, err.Error())
	}
	switch data.Type {
	case "friend":
		res, err := user_record.GetListAndTotal(&user_record.GetListRequest{
			UserID:   c.GetUserID(r),
			FriendID: data.ID,
			Page:     data.Page,
			Limit:    data.Limit,
		})
		if err != nil {
			c.Fail(r, err.Error())
		}
		c.Success(r, res)
	case "group":
		res, err := group_record.GetListAndTotal(&group_record.GetListRequest{
			GroupID: data.ID,
			Page:    data.Page,
			Limit:   data.Limit,
		})
		if err != nil {
			c.Fail(r, err.Error())
		}
		c.Success(r, res)
	default:
		c.Fail(r, "类型不正确")

	}

}
