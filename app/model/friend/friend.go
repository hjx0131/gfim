package friend

import (
	"github.com/gogf/gf/container/gvar"
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
		Fields("f.*,u.nickname,u.avatar,u.im_status,u.sign").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}

//GetFriendUserIds 获取好友id列表
func GetFriendUserIds(UserID uint) ([]*gvar.Var, error) {
	list, err := g.DB().
		Table(Table).
		Array("friend_id", "user_id=?", UserID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
