package friend

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetListByUserID 获取好友列表
func GetListByUserID(UserID uint) (gdb.Result, error) {
	list, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.friend_id").
		Where("f.user_id=?", UserID).
		Fields("f.*,u.nickname,u.avatar,u.status,u.bio").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
