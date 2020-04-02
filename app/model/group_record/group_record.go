package group_record

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetListAndTotal 获取好友聊天记录和聊天总数
func GetListAndTotal(GroupID uint, Page, Limit int) (gdb.Result, int, error) {
	list, err := g.DB().
		Table(Table).
		As("g").
		InnerJoin("gf_user u", "u.id=g.user_id").
		Where("g.group_id=?", GroupID).
		Fields("g.user_id,g.content,g.createtime,u.nickname,u.avatar").
		Page(Page, Limit).
		All()
	if err != nil {
		return nil, 0, err
	}
	count, err := g.DB().
		Table(Table).
		As("g").
		InnerJoin("gf_user u", "u.id=g.user_id").
		Where("g.group_id=?", GroupID).
		Count()
	if err != nil {
		return list, 0, err
	}
	return list, count, nil
}
