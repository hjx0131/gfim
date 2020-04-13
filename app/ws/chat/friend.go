package chat

import (
	"fmt"
	"gfim/app/model/friend"
	"gfim/app/model/user"
	"gfim/app/model/user_record"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	f := userIds.Get(freq.ToUserID)
	if f != nil {
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
		writeByWs(f.(*ghttp.WebSocket), resp)
		// .WriteMessage(ghttp.WS_MSG_TEXT, data)
	}
	return nil
}

//writeFriends 发送消息给所有好友,状态切换通知等
func (c *Controller) writeFriends(userID uint, resp *MsgResp) error {
	ids, err := friend.GetFriendUserIds(userID)
	if err != nil {
		return err
	}
	if ids == nil {
		fmt.Println("没有可通知的好友")
		return nil
	}
	var toUserID uint
	for _, id := range ids {
		toUserID = id.Uint()
		f := userIds.Get(toUserID)
		if f != nil {
			writeByWs(f.(*ghttp.WebSocket), resp)

			// f.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, data)
		}
	}
	return nil
}
