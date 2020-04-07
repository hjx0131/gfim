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
func GetListByUserID(userID uint) ([]*GroupInfo, error) {
	//好友群组列表
	glist, err := friend_group.GetListByUserID(userID)
	if err != nil {
		return nil, err
	}
	//好友列表
	flist, _ := friend.GetListByUserID(userID)
	res := make([]*GroupInfo, len(glist))
	if glist != nil {
		for index, item := range glist {
			f := make([]*FriendInfo, 0)
			if flist != nil {
				for _, val := range flist {
					if val["friend_group_id"].Uint() == item.Id {
						imStatus := val["im_status"].String()
						//隐身状态显示离线
						if imStatus == "hide" {
							imStatus = "offline"
						}
						f = append(f, &FriendInfo{
							ID:       val["friend_id"].Uint(),
							Username: val["nickname"].String(),
							Avatar:   val["avatar"].String(),
							Sign:     val["sign"].String(),
							Status:   imStatus,
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
