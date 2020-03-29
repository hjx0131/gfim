package chat

import (
	"gfim/app/model/user"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

//FriendReq 好友聊天请求格式
type FriendReq struct {
	FormUserID uint   `json:"from_user_id" v:"from_user_id@required#发送人不能为空`
	ToUserID   uint   `json:"to_user_id" v:"to_user_id@required#接收人不能为空`
	Content    string `json:"content" v:"content@required#内容不能为空`
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
	//数据添加到数据库...

	//如果接收人在线， 发送消息
	f := users.Get(freq.ToUserID)
	if f != nil {
		one, err := user.FindOne("id=?", freq.FormUserID)
		if err != nil {
			return err
		}
		now := gtime.Timestamp()
		resp := &FriendResp{
			Username:  one.Nickname,
			Avatar:    one.Avatar,
			ID:        one.Id,
			Type:      "friend",
			Content:   freq.Content,
			Cid:       1,
			Mine:      false,
			FromID:    one.Id,
			Timestamp: gconv.Uint(now) * 1000,
		}
		data, err := gjson.Encode(resp)
		if err != nil {
			return err
		}
		f.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, data)
	}
	return nil
}
