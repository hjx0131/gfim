package apply

import (
	"errors"
	"gfim/app/model/apply"
	"gfim/app/model/user"

	"gfim/app/model/apply_remind"
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
	GroupID  uint `json:"friend_group_id" v:"group_id@required#群组不能为空"`
	ToUserID uint
	Remark   string `json:"remark"`
}

//HandleReq 处理申请请求参数
type HandleReq struct {
	ID            uint `json:"id" v:"id@required#申请id不能为空"`
	FriendGroupID uint `json:"friend_group_id"`
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
	res, err := apply.Model.Data(g.Map{
		"type":         "friend",
		"from_user_id": userID,
		"to_user_id":   req.FriendID,
		"target_id":    req.FriendGroupID,
		"remark":       req.Remark,
		"create_time":  now,
	}).Insert()
	if err != nil {
		return err
	}
	//向接收人添加一条验证提醒
	applyID, _ := res.LastInsertId()
	apply_remind.Model.Data(g.Map{
		"apply_id":    applyID,
		"user_id":     req.FriendID,
		"create_time": now,
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
	res, err := apply.Model.Data(g.Map{
		"type":         "group",
		"from_user_id": userID,
		"to_user_id":   group.UserId,
		"target_id":    req.GroupID,
		"remark":       req.Remark,
		"create_time":  now,
	}).Insert()
	req.ToUserID = group.UserId
	//向接收人添加一条验证提醒
	applyID, _ := res.LastInsertId()
	apply_remind.Model.Data(g.Map{
		"apply_id":    applyID,
		"user_id":     group.UserId,
		"create_time": now,
	}).Insert()
	return nil
}

//HandleCheck 1
func HandleCheck(handleUserID uint, req *HandleReq) (*apply.Entity, error) {
	// 数据校验
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return nil, err
	}
	one, err := apply.Model.Where("id", req.ID).FindOne()
	if err != nil {
		return one, err
	}
	if one == nil {
		return nil, errors.New("申请未找到")
	}
	if one.State != 1 {
		return one, errors.New("当前状态下不能同意")
	}
	if handleUserID != one.ToUserId {
		return one, errors.New("您没有处理权限")
	}
	return one, nil
}

//Agree 同意
func Agree(handleUserID uint, req *HandleReq) (*apply.Entity, error) {
	one, err := HandleCheck(handleUserID, req)
	if err != nil {
		return one, err
	}
	// if req.FriendGroupID <= 0 {
	// 	return one, errors.New("好友分组不能为空")
	// }
	now := gtime.Timestamp()
	apply.Model.
		Where("id", one.Id).
		Data(g.Map{
			"handle_time": now,
			"state":       2,
		}).
		Update()
	if one.Type == "friend" {
		//建立好友关联
		friend.Model.Data(g.Map{
			"user_id":         one.FromUserId,
			"friend_id":       one.ToUserId,
			"friend_group_id": one.TargetId,
			"create_time":     now,
		}).Insert()
		friend.Model.Data(g.Map{
			"user_id":         one.ToUserId,
			"friend_id":       one.FromUserId,
			"friend_group_id": req.FriendGroupID,
			"create_time":     now,
		}).Insert()
	} else {
		//建立群和用户关联
		group_user.Model.Data(g.Map{
			"group_id":    one.TargetId,
			"user_id":     one.FromUserId,
			"create_time": now,
		}).Insert()
	}
	//向发起人添加一条同意提醒
	apply_remind.Model.Data(g.Map{
		"apply_id":    one.Id,
		"user_id":     one.FromUserId,
		"create_time": now,
	}).Insert()
	return one, nil
}

//Refuse 拒绝
func Refuse(handleUserID uint, req *HandleReq) (*apply.Entity, error) {
	one, err := HandleCheck(handleUserID, req)
	if err != nil {
		return one, err
	}
	now := gtime.Timestamp()
	apply.Model.
		Where("id", one.Id).
		Data(g.Map{
			"handle_time": now,
			"state":       3,
		}).
		Update()
	//向发起人添加一条拒绝提醒
	apply_remind.Model.Data(g.Map{
		"apply_id":    one.Id,
		"user_id":     one.FromUserId,
		"create_time": now,
	}).Insert()
	return one, nil
}

//GetListRequest 获取记录所需要的参数
type GetListRequest struct {
	UserID uint
	Page   int
	Limit  int
}

//User  用户信息格式
type User struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	Avatar string `json:"avatar"`
}

//Info 申请消息格式
type Info struct {
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
	GroupName string `json:"groupname"`
	Remark    string `json:"remark"`
	TargetID  uint   `json:"group_id"`
	State     uint   `json:"state"`
	StateText string `json:"state_text"`
	FromSelf  bool   `json:"from_self"`
	User      *User  `json:"user"`
	CanHandle bool   `json:"can_handle"`
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
			var canHandle bool
			var groupname string
			//好友验证
			if item.Type == "friend" {
				//发起人是自己
				if req.UserID == item.FromUserId {
					fromSelf = true
					uid = item.ToUserId
					content = "验证消息已发送"
				} else {
					if item.State == 1 {
						canHandle = true
					}
					fromSelf = false
					uid = item.FromUserId
					content = "申请添加你为好友"
				}

			} else {
				one, _ := group.Model.Where("id", item.TargetId).One()
				if one != nil {
					groupname = one.Name
				}
				//群验证
				if req.UserID == item.FromUserId {
					//发起人是自己
					fromSelf = true
					uid = item.ToUserId
					content = "验证消息已发送"
				} else {
					if item.State == 1 {
						canHandle = true
					}
					fromSelf = false
					uid = item.FromUserId
					content = "申请进群"

				}
			}
			one, _ := user.Model.Where("id", uid).FindOne()
			res[index] = &Info{
				ID:        item.Id,
				Type:      item.Type,
				Content:   content,
				Timestamp: gconv.Int(item.CreateTime) * 1000,
				GroupName: groupname,
				Remark:    item.Remark,
				TargetID:  item.TargetId,
				State:     item.State,
				StateText: item.GetStateText(),
				FromSelf:  fromSelf,
				User: &User{
					ID:     one.Id,
					Name:   one.Nickname,
					Avatar: one.Avatar,
				},
				CanHandle: canHandle,
			}
		}
	}
	data := g.Map{
		"list":  res,
		"count": count,
	}
	return data, nil
}
