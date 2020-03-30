package chat

import (
	"fmt"
	"gfim/app/model/friend"
	"gfim/app/model/user"
	"gfim/app/service/user_token"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var (
	//uses 用户
	users = gmap.New(true)
)

//MsgReq 接收消息结构体
type MsgReq struct {
	Type  string      `json:"type" v:"type@required#消息类型不能为空"`
	Data  interface{} `json:"data" v:""`
	Token string      `json:"token" v:"token@required#token不能为空"`
}

//MsgResp 发送消息结构体
type MsgResp struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

//Controller 控制器
type Controller struct {
	ws *ghttp.WebSocket
}

//WebSocket ws
func (c *Controller) WebSocket(r *ghttp.Request) {
	msg := &MsgReq{}

	// 初始化WebSocket请求
	ws, err := r.WebSocket()
	if err != nil {
		user.Model.Data("im_status", "offline").Update()
		glog.Error(err)
		r.Exit()
	}
	c.ws = ws

	for {
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("没了")
			return
		}
		// json解析
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			c.write(MsgResp{"error", "消息格式不正确: " + err.Error()})
			continue
		}
		// 数据校验
		if err := gvalid.CheckStruct(msg, nil); err != nil {
			c.write(MsgResp{"error", err.String()})
			continue
		}
		// 检验token
		UserID, err := getUserID(msg.Token)
		if err != nil {
			c.write(MsgResp{"error", err.Error()})
			continue
		}
		//发送消息
		switch msg.Type {
		case "confirmJoin": //加入连接
			c.joinConn(UserID)
		case "close": //退出连接
			c.closeConn(UserID)
		case "ping": //心跳检测
			c.write(MsgResp{"ping", "pong"})
			c.write(MsgResp{"welcome", "欢迎" + gconv.String(UserID)})
		case "friend": //好友聊天
			err := c.FriendChat(msg)
			if err != nil {
				c.write(MsgResp{"error", err.Error()})
				continue
			}
		}

	}
}
func (c *Controller) write(msg MsgResp) error {
	data, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return c.ws.WriteMessage(ghttp.WS_MSG_TEXT, data)
}

//getUserID 获取UserID
func getUserID(token string) (uint, error) {
	data := &user_token.GetIDInput{
		Token: token,
	}
	UserID, e := user_token.GetUserID(data)
	if e != nil {
		return 0, e
	}
	return UserID, nil
}

//closeConn 断开连接处理
func (c *Controller) closeConn(UserID uint) {
	users.Remove(UserID)
	//状态修改为下线，并通知好友
	user.Model.Data("im_status", "offline").Where("id=?", UserID).Update()
	err := c.writeFriends(UserID, &MsgResp{
		Type: "offline",
		Data: UserID,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//joinConn 加入连接处理
func (c *Controller) joinConn(UserID uint) {
	users.Set(UserID, c.ws)
	//状态修改为在线，并通知好友
	user.Model.Data("im_status", "online").Where("id=?", UserID).Update()
	err := c.writeFriends(UserID, &MsgResp{
		Type: "online",
		Data: UserID,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//writeFriends 发送消息给所有好友,状态切换通知等
func (c *Controller) writeFriends(UserID uint, resp *MsgResp) error {
	data, err := gjson.Encode(resp)
	if err != nil {
		return err
	}
	ids, err := friend.GetFriendUserIds(UserID)
	if err != nil {
		return err
	}
	if ids == nil {
		fmt.Println("没有可通知的好友")
		return nil
	}
	fmt.Printf("%#v", users)
	var toUserID uint
	for _, id := range ids {
		toUserID = id.Uint()
		f := users.Get(toUserID)
		if f != nil {
			f.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, data)
		}
	}
	return nil
}
