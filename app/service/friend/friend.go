package friend

import (
	"gfim/app/model/friend"
)

//UserFriendList 用户好友列表
func UserFriendList(userID uint) (interface{}, error) {
	list, err := friend.GetListByUserID(userID)
	return list, err
}
