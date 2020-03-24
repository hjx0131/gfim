package friend_group

import (
	"gfim/app/model/friend"
	"gfim/app/model/friend_group"
)

//GroupInfo 好友群组结构体,前台需要的格式
type GroupInfo struct {
	ID        uint          `json:"id"`
	GroupName string        `json:"groupname"`
	List      []*FriendInfo `json:"list"`
}

//FriendInfo 好友结构体，前台需要的格式
type FriendInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Status   string `json:"status"`
}

// GetListByUserID 根据用户ID获取好友分组列表
func GetListByUserID(UserID uint) ([]*GroupInfo, error) {
	list, err := friend_group.GetListByUserID(UserID)
	if err != nil {
		return nil, err
	}
	flist, _ := friend.GetListByUserID(UserID)

	res := make([]*GroupInfo, len(list))

	if list != nil {
		for index, item := range list {
			f := make([]*FriendInfo, 0)
			if flist != nil {
				for _, val := range flist {
					if val.FriendGroupId == item.Id {
						f = append(f, &FriendInfo{
							ID: val.Id,
							// Username: val.Nickname,
							// Avatar:   val.Avatar,
							// Sign:     val.Bio,
							// Status:   val.Status,
						})
					}
				}
			}
			res[index] = &GroupInfo{
				ID:        item.Id,
				GroupName: item.Name,
				List:      f,
			}
		}
	}
	return res, nil
}
