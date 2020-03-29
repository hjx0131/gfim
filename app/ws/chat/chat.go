package chat

import (
	"gfim/app/service/user_token"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var (
	//usersConn 用户连接
	usersConn = gmap.New(true)
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
		glog.Error(err)
		r.Exit()
	}
	c.ws = ws
	for {
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			users.Remove(usersConn.Get(ws))
			usersConn.Remove(ws)
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
		usersConn.Set(ws, UserID)
		users.Set(UserID, ws)
		//发送消息
		switch msg.Type {
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
