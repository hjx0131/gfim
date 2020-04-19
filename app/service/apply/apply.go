package apply

import (
	"errors"
	"gfim/app/model/apply"
	"gfim/app/model/user"

	"gfim/app/model/friend"
	"gfim/app/model/group"
	"gfim/app/model/group_user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

//FriendReq 好友申请请求参数
type FriendReq struct {
	FriendID      uint   `json:"friend_id" v:"friend_id@required#好友不能为空"`
	FriendGroupID uint   `json:"friend_group_id" v:"friend_group_id@required#好友分组不能为空"`
	Remark        string `json:"remark"`
}

//GroupReq 群组申请请求参数
type GroupReq struct {
	GroupID uint   `json:"friend_group_id" v:"friend_group_id@required#群组不能为空"`
	Remark  string `json:"remark"`
}

//Friend 好友申请
func Friend(userID uint, req *FriendReq) error {
	// 数据校验
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return err
	}
	count, err := user.Model.
		Where("id", req.FriendID).
		Count()
	if count == 0 {
		return errors.New("用户不存在")
	}
	//不能添加自己和好友，不能重复申请
	if userID == req.FriendID {
		return errors.New("您不能添加自己")
	}
	count, err = friend.Model.
		Where("user_id", userID).
		Where("friend_id", req.FriendID).
		Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("对方已经是您的好友")
	}
	count, err = apply.Model.
		Where("from_user_id", userID).
		Where("to_user_id", req.FriendID).
		Where("type", "friend").
		Where("state", 1).
		Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("等待对方验证中，请勿重复提交")
	}
	now := gtime.Timestamp()
	apply.Model.Data(g.Map{
		"type":         "friend",
		"from_user_id": userID,
		"to_user_id":   req.FriendID,
		"target_id":    req.FriendGroupID,
		"remark":       req.Remark,
		"create_time":  now,
	}).Insert()
	return nil
}

//Group  群组申请
func Group(userID uint, req *GroupReq) error {
	// 数据校验
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return err
	}
	// 检查群组是否存在
	group, err := group.Model.Where("id", req.GroupID).One()
	if err != nil {
		return err
	}
	if group == nil {
		return errors.New("群组不存在")
	}
	//不能添加已加入的群组，不能重复申请
	count, err := group_user.Model.
		Where("user_id", userID).
		Where("group_id", req.GroupID).
		Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("您已经在该群中")
	}
	count, err = apply.Model.
		Where("from_user_id", userID).
		Where("target_id", req.GroupID).
		Where("state", 1).
		Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("等待对方验证中，请勿重复提交")
	}
	now := gtime.Timestamp()
	apply.Model.Data(g.Map{
		"type":         "friend",
		"from_user_id": userID,
		"to_user_id":   group.UserId,
		"target_id":    req.GroupID,
		"remark":       req.Remark,
		"create_time":  now,
	}).Insert()
	return nil
}

func Agree() {

}
func Refuse() {

}

//GetListRequest 获取记录所需要的参数
type GetListRequest struct {
	UserID uint
	Page   int
	Limit  int
}

//UserInfo  用户信息格式
type UserInfo struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
}

//Info 申请消息格式
type Info struct {
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
	Remark    string `json:"remark"`
	State     uint   `json:"state"`
	FromSelf  bool   `json:"from_self"`
	UserInfo  *UserInfo
}

//GetListAndTotal 获取列表和数量
func GetListAndTotal(req *GetListRequest) (interface{}, error) {
	list, count, err := apply.GetListAndTotal(req.UserID, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]*Info, len(list))
	if list != nil {
		for index, item := range list {
			var uid uint
			var content string
			var fromSelf bool
			//发起人是自己
			if req.UserID == item.FromUserId {
				fromSelf = true
				uid = item.ToUserId
				content = "验证已发送"
			} else {
				fromSelf = false
				uid = item.FromUserId
				content = "申请添加你为好友"
			}
			one, _ := user.Model.Where("id", uid).FindOne()
			res[index] = &Info{
				ID:        item.Id,
				Timestamp: gconv.Int(item.CreateTime) * 1000,
				Content:   content,
				Type:      item.Type,
				FromSelf:  fromSelf,
				State:     item.State,
				Remark:    item.Remark,
				UserInfo: &UserInfo{
					ID:       one.Id,
					Username: one.Nickname,
					Avatar:   one.Avatar,
				},
			}
		}
	}
	data := g.Map{
		"list":  res,
		"count": count,
	}
	return data, nil
}
