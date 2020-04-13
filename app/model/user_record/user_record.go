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
		Fields("f.user_id,f.content,f.create_time,u.nickname,u.avatar").
		Order("f.id desc").
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

//GetNoNotifyRecord 获取还未通知的消息
func GetNoNotifyRecord(userID uint) (gdb.Result, error) {
	list, err := g.DB().
		Table(Table).
		As("r").
		InnerJoin("gf_user u", "u.id=r.user_id").
		Where("r.friend_id=?", userID).
		Where("r.is_notify", 0).
		Fields("r.id,r.user_id,r.content,r.create_time,u.nickname,u.avatar").
		Order("r.id desc").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
