package group

import (
	"gfim/app/model/group"
)

//UserInfo 群员结构体,前台需要的格式
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
}

//Info 群结构体,前台需要的格式
type Info struct {
	ID        uint   `json:"id"`
	GroupName string `json:"groupname"`
	Avatar    string `json:"avatar"`
}

// GetUserListByID 根据群ID获取群员列表
func GetUserListByID(ID uint) ([]*UserInfo, error) {
	list, err := group.GetUserListByID(ID)
	if err != nil {
		return nil, err
	}
	res := make([]*UserInfo, len(list))
	if list != nil {
		for index, item := range list {
			res[index] = &UserInfo{
				ID:       item["id"].Uint(),
				Username: item["nickname"].String(),
				Avatar:   item["avatar"].String(),
				Sign:     item["sign"].String(),
			}
		}
	}
	return res, nil
}

//GetListByUserID 根据用户ID获取群列表
func GetListByUserID(UserID uint) ([]*Info, error) {
	list, err := group.Model.As("g").
		Where("u.user_id=?", UserID).
		InnerJoin("gf_group_user u", "u.group_id=g.id").
		Fields("g.id,g.name,g.avatar").
		All()
	if err != nil {
		return nil, err
	}
	res := make([]*Info, len(list))

	if list != nil {
		for index, item := range list {
			res[index] = &Info{
				ID:        item.Id,
				GroupName: item.Name,
				Avatar:    item.Avatar,
			}
		}
	}
	return res, nil
}
