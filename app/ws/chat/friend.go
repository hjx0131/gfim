package chat

import (
	"gfim/app/model/friend"
	"gfim/app/model/user"
	"gfim/app/model/user_record"
	"gfim/app/service/apply"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

//FriendReq 好友聊天请求格式
type FriendReq struct {
	FormUserID uint   `json:"from_user_id" v:"from_user_id@required#发送人不能为空"`
	ToUserID   uint   `json:"to_user_id" v:"to_user_id@required#接收人不能为空"`
	Content    string `json:"content" v:"content@required#内容不能为空"`
}

// FriendResp 好友聊天返回格式
type FriendResp struct {
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Cid       uint   `json:"cid"`
	Mine      bool   `json:"mine"`
	FromID    uint   `json:"fromid"`
	Timestamp uint   `json:"timestamp"`
}

//FriendChat 好友聊天
func (c *Controller) FriendChat(msg *MsgReq) error {
	freq := &FriendReq{}
	err := gconv.Struct(msg.Data, freq)
	if err != nil {
		panic(err)
	}
	// 数据校验
	if err := gvalid.CheckStruct(freq, nil); err != nil {
		return err
	}
	now := gtime.Timestamp()
	//数据添加到数据库...
	res, err := user_record.Model.Insert(g.Map{
		"user_id":     freq.FormUserID,
		"friend_id":   freq.ToUserID,
		"content":     freq.Content,
		"create_time": now,
	})
	if err != nil {
		return err
	}
	recordID, _ := res.LastInsertId()
	//如果接收人在线， 发送消息
	ws := getWsByUserID(freq.ToUserID)
	if ws != nil {
		one, err := user.FindOne("id=?", freq.FormUserID)
		if err != nil {
			return err
		}
		resp := &MsgResp{
			Type: "friend",
			Data: &FriendResp{
				Username:  one.Nickname,
				Avatar:    one.Avatar,
				ID:        one.Id,
				Type:      "friend",
				Content:   freq.Content,
				Cid:       gconv.Uint(recordID),
				Mine:      false,
				FromID:    one.Id,
				Timestamp: gconv.Uint(now) * 1000,
			},
		}
		writeByWs(ws, resp)
		//修改为已通知
		user_record.Model.Where("id=?", recordID).
			Data("is_notify", 1).
			Update()
	}
	return nil
}

//writeFriends 发送统一的消息给所有好友,状态切换通知等
func (c *Controller) writeFriends(userID uint, resp *MsgResp) error {
	ids, err := friend.GetFriendUserIds(userID)
	if err != nil {
		return err
	}
	if ids == nil {
		return nil
	}
	var toUserID uint
	for _, id := range ids {
		toUserID = id.Uint()
		ws := getWsByUserID(toUserID)
		if ws != nil {
			writeByWs(ws, resp)
		}
	}
	return nil
}

//notifyUserRecord 推送好友留言(未通知的聊天记录)
func (c *Controller) notifyUserRecord(userID uint) error {
	list, err := user_record.GetNoNotifyRecord(userID)
	if err != nil {
		return err
	}
	data := make([]*FriendResp, len(list))
	if list != nil {
		for index, item := range list {
			data[index] = &FriendResp{
				Username:  item["nickname"].String(),
				Avatar:    item["avatar"].String(),
				ID:        item["user_id"].Uint(),
				Type:      "friend",
				Content:   item["content"].String(),
				Cid:       item["id"].Uint(),
				Mine:      false,
				FromID:    item["user_id"].Uint(),
				Timestamp: item["create_time"].Uint() * 1000,
			}
		}
		resp := &MsgResp{
			Type: NotifyRecord,
			Data: data,
		}
		c.writeByUserID(userID, resp)
		//修改为已通知
		user_record.Model.
			Where("friend_id=?", userID).
			Where("is_notify", 0).
			Data("is_notify", 1).
			Update()
	}
	return nil
}

type applyReq struct {
	friend apply.FriendReq
	group  apply.GroupReq
}

//applyFriend 好友申请
func (c *Controller) apply(userID uint, msg *MsgReq) error {
	req := &applyReq{}
	err := gconv.Struct(msg.Data, &req.friend)
	if err != nil {
		panic(err)
	}
	if err := apply.Friend(userID, &req.friend); err != nil {
		return err
	}
	//向该好友推送未读的好友申请
	c.NoReadApplyCount(req.friend.FriendID)
	return nil
}

type handleReq struct {
	friend apply.HandleReq
}

//agree 同意好友申请
func (c *Controller) agree(userID uint, msg *MsgReq) error {
	req := &handleReq{}
	err := gconv.Struct(msg.Data, &req.friend)
	if err != nil {
		panic(err)
	}
	one, err := apply.Agree(userID, &req.friend)
	if err != nil {
		return err
	}
	//向发起人推送未读提醒
	c.NoReadApplyCount(one.FromUserId)
	//追加好友到面板
	c.AppendFriend(userID, one.FromUserId)
	ws := getWsByUserID(one.FromUserId)
	if ws != nil {
		c.AppendFriend(one.FromUserId, userID)
	}
	return nil
}

//refuse 拒绝好友申请
func (c *Controller) refuse(userID uint, msg *MsgReq) error {
	req := &handleReq{}
	err := gconv.Struct(msg.Data, &req.friend)
	if err != nil {
		panic(err)
	}
	one, err := apply.Refuse(userID, &req.friend)
	if err != nil {
		return nil
	}
	//向发起人推送
	c.NoReadApplyCount(one.FromUserId)
	return nil
}

//AppendFriend 追加好友到面板
func (c *Controller) AppendFriend(userID, friendID uint) error {
	friendInfo, err := g.DB().
		Table(friend.Table).
		As("f").
		InnerJoin("gf_user u", "u.id=f.friend_id").
		Where(g.Map{
			"f.user_id":   userID,
			"f.friend_id": friendID,
		}).
		Fields("f.friend_group_id,u.id,u.nickname,u.avatar,u.im_status,u.sign").
		One()
	if err != nil {
		return err
	}
	if friendInfo != nil {
		c.writeByUserID(userID, &MsgResp{
			Type: AppendFriend,
			Data: g.Map{
				"type":     "friend",
				"avatar":   friendInfo["avatar"].String(),
				"username": friendInfo["nickname"].String(),
				"groupid":  friendInfo["friend_group_id"].Int(),
				"id":       friendInfo["id"].Int(),
				"sign":     friendInfo["sign"].String(),
			},
		})
	}
	return nil
}
