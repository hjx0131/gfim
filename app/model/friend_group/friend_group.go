package friend_group

import (
	"github.com/gogf/gf/frame/g"
)

// GetListByUserID 根据用户ID获取好友分组列表
func GetListByUserID(UserID uint) ([]*Entity, error) {
	list, e := Model.
		Where("user_id in(?)", g.Slice{0, UserID}).
		Fields("id,name").
		FindAll()
	if e != nil {
		return nil, e
	}
	return list, e
}
