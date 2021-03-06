package friend

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//GetListByUserID 获取好友列表
func GetListByUserID(userID uint) (gdb.Result, error) {
	list, err := g.DB().
		Table(Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.friend_id").
		Where("f.user_id=?", userID).
		Fields("f.*,u.nickname,u.avatar,u.im_status,u.sign").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}

//GetFriendUserIds 获取好友id列表
func GetFriendUserIds(userID uint) ([]*gvar.Var, error) {
	list, err := g.DB().
		Table(Table).
		Array("friend_id", "user_id=?", userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//GetTwoUserInfo 获取互为好友的用户信息
func GetTwoUserInfo(userID, friendID uint) (gdb.Result, error) {
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
		Fields("f.friend_group_id,u.id,u.nickname,u.avatar,u.im_status,u.sign").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
