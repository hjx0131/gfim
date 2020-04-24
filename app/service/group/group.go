package group

import (
	"gfim/app/model/group"
	"gfim/app/model/group_user"

	"gfim/library/avatar"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
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
func GetListByUserID(userID uint) ([]*Info, error) {
	list, err := group.Model.As("g").
		Where("u.user_id=?", userID).
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

//SaveReq 保存请求参数
type SaveReq struct {
	UserID uint   `json:"user_id" v:"user_id@required#创建人不能为空"`
	Name   string `json:"name" v:"name@required#群名不能为空"`
}

//Save 保存数据
func Save(req *SaveReq) error {
	// 数据校验
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return err
	}
	now := gtime.Timestamp()
	res, err := group.Model.
		Data(g.Map{
			"user_id":        req.UserID,
			"create_user_id": req.UserID,
			"name":           req.Name,
			"create_time":    now,
			"avatar":         avatar.NewRandom(req.Name),
		}).
		Insert()
	if err != nil {
		return err
	}
	groupID, _ := res.LastInsertId()
	group_user.Model.
		Data(g.Map{
			"user_id":     req.UserID,
			"group_id":    groupID,
			"create_time": now,
		}).
		Insert()
	return nil
}

//SearchReq 搜索列表参数
type SearchReq struct {
	Wd    string `json:"wd"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

//SearchResp 搜索返回
type SearchResp struct {
	List  []*Info `json:"list"`
	Count int     `json:"count"`
}

//Search 查找群组
func Search(req *SearchReq) (*SearchResp, error) {
	where := make(map[string]string)
	if req.Wd != "" {
		where["name like ?"] = "%" + req.Wd + "%"
	}
	all, err := group.Model.
		Where(where).
		Order("id desc").
		Page(req.Page, req.Limit).
		All()
	if err != nil {
		return nil, err
	}
	count, err := group.Model.
		Where(where).
		Count()
	if err != nil {
		return nil, err
	}
	list := make([]*Info, len(all))
	if all != nil {
		for index, item := range all {
			list[index] = &Info{
				ID:        item.Id,
				GroupName: item.Name,
				Avatar:    item.Avatar,
			}
		}
	}
	res := &SearchResp{
		List:  list,
		Count: count,
	}
	return res, nil
}
