package group

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// GetUserListByID 根据群ID获取群员列表
func GetUserListByID(ID uint) (gdb.Result, error) {
	list, err := g.DB().
		Table("gf_group_user").
		As("g").
		InnerJoin("gf_user u", "u.id=g.user_id").
		Where("g.group_id=?", ID).
		Fields("g.id,g.nickname as gnickname,u.nickname,u.avatar,u.status,u.bio").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
