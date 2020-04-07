package user_record

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetListAndTotal 获取好友聊天记录和聊天总数
func GetListAndTotal(userID, friendID uint, page, limit int) (gdb.Result, int, error) {
	list, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.user_id").
		Where(g.Map{
			"f.user_id":   userID,
			"f.friend_id": friendID,
		}).
		Or(g.Map{
			"f.user_id":   friendID,
			"f.friend_id": userID,
		}).
		Fields("f.user_id,f.content,f.createtime,u.nickname,u.avatar").
		Page(page, limit).
		All()
	if err != nil {
		return nil, 0, err
	}
	count, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.user_id").
		Where(g.Map{
			"f.user_id": userID,
			"friend_id": friendID,
		}).
		Or(g.Map{
			"f.user_id":   friendID,
			"f.friend_id": userID,
		}).
		Count()
	if err != nil {
		return list, 0, err
	}
	return list, count, nil
}
