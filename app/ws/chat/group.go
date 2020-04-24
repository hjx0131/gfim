package chat

import (
	"gfim/app/model/group_record"
	"gfim/app/model/group_user"
	"gfim/app/model/user"
	"gfim/app/service/apply"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

//GroupReq 群聊天请求格式
type GroupReq struct {
	UserID  uint   `json:"user_id" v:"user_id@required#发送人不能为空"`
	GroupID uint   `json:"group_id" v:"group_id@required#群不能为空"`
	Content string `json:"content" v:"content@required#内容不能为空"`
}

// GroupResp 群聊天返回格式
type GroupResp struct {
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

//GroupChat 群聊天
func (c *Controller) GroupChat(msg *MsgReq) error {
	freq := &GroupReq{}
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
	res, err := group_record.Model.Insert(g.Map{
		"user_id":     freq.UserID,
		"group_id":    freq.GroupID,
		"content":     freq.Content,
		"create_time": now,
	})
	if err != nil {
		return err
	}
	one, err := user.FindOne("id=?", freq.UserID)
	if err != nil {
		return err
	}
	recordID, _ := res.LastInsertId()
	//获取所有群员
	list, err := group_user.GetGroupUserList(freq.GroupID)
	if err != nil {
		return err
	}
	if list != nil {
		for _, item := range list {
			//如果在线，则发送消息，不给自己推送.
			f := userIds.Get(item["user_id"].Uint())
			if f != nil && item["user_id"].Uint() != freq.UserID {
				resp := &MsgResp{
					Type: "group",
					Data: &FriendResp{
						Username:  one.Nickname,
						Avatar:    one.Avatar,
						ID:        freq.GroupID,
						Type:      "group",
						Content:   freq.Content,
						Cid:       gconv.Uint(recordID),
						Mine:      false,
						FromID:    freq.UserID,
						Timestamp: gconv.Uint(now) * 1000,
					},
				}
				writeByWs(f.(*ghttp.WebSocket), resp)
			}

		}
	}
	return nil
}

//applyGroup 群组申请
func (c *Controller) applyGroup(userID uint, msg *MsgReq) error {
	req := &applyReq{}
	err := gconv.Struct(msg.Data, &req.group)
	if err != nil {
		panic(err)
	}
	if err := apply.Group(userID, &req.group); err != nil {
		return err
	}
	//向该好友推送未读的好友申请
	c.NoReadApplyCount(req.group.ToUserID)
	return nil
}
