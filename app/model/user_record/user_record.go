package user_record

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetListAndTotal 获取好友聊天记录和聊天总数
func GetListAndTotal(UserID, FriendID uint, Page, Limit int) (gdb.Result, int, error) {
	list, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.user_id").
		Where(g.Map{
			"f.user_id":   UserID,
			"f.friend_id": FriendID,
		}).
		Or(g.Map{
			"f.user_id":   FriendID,
			"f.friend_id": UserID,
		}).
		Fields("f.user_id,f.content,f.createtime,u.nickname,u.avatar").
		Page(Page, Limit).
		All()
	if err != nil {
		return nil, 0, err
	}
	count, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.user_id").
		Where(g.Map{
			"f.user_id": UserID,
			"friend_id": FriendID,
		}).
		Or(g.Map{
			"f.user_id":   FriendID,
			"f.friend_id": UserID,
		}).
		Count()
	if err != nil {
		return list, 0, err
	}
	return list, count, nil
}
