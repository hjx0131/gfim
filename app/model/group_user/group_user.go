package group_user

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetGroupUserList 获取群成员
func GetGroupUserList(groupID uint) (gdb.Result, error) {
	list, err := g.DB().
		Table(Table).
		As("g").
		InnerJoin("gf_user u", "u.id=g.user_id").
		Where("g.group_id=?", groupID).
		Fields("g.user_id,g.group_id,g.create_time,u.nickname,u.avatar").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
