package group_record

import (
	"gfim/app/model/group_record"

	"github.com/gogf/gf/frame/g"
)

//GetListRequest 获取记录所需要的参数
type GetListRequest struct {
	GroupID uint
	Page    int
	Limit   int
}

//Info 1
type Info struct {
	Username  string `json:"username"`
	ID        uint   `json:"id"`
	Avatar    string `json:"avatar"`
	Timestamp int    `json:"timestamp"`
	Content   string `json:"content"`
}

//GetListAndTotal 获取好友聊天记录和聊天总数
func GetListAndTotal(req *GetListRequest) (interface{}, error) {
	list, count, err := group_record.GetListAndTotal(req.GroupID, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]*Info, len(list))
	if list != nil {
		for index, item := range list {
			res[index] = &Info{
				ID:        item["user_id"].Uint(),
				Username:  item["nickname"].String(),
				Avatar:    item["avatar"].String(),
				Timestamp: item["createtime"].Int() * 1000,
				Content:   item["content"].String(),
			}
		}
	}
	data := g.Map{
		"list":  res,
		"count": count,
	}
	return data, nil
}
