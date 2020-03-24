package friend

import (
	"gfim/app/model/friend"
)

//UserFriendList 用户好友列表
func UserFriendList(UserID uint) (interface{}, error) {
	list, err := friend.GetListByUserID(UserID)
	return list, err
}
