package chat

import (
	"fmt"
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
	users   = gmap.New(true) //存储{ws:userID}
	userIds = gmap.New(true) //存储{userID:ws}
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
			//断开连接处理
			c.closeConn(ws)
			fmt.Println("断开连接")
			break
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
		userID, err := getUserID(msg.Token)
		if err != nil {
			c.write(MsgResp{"error", err.Error()})
			continue
		}
		//发送消息
		switch msg.Type {
		case "confirmJoin": //确认连接
			c.joinConn(userID)
		case "close": //退出连接
		case "ping": //心跳检测
			c.write(MsgResp{"ping", "pong"})
			c.write(MsgResp{"welcome", "欢迎" + gconv.String(userID)})
		case "friend": //好友聊天
			err := c.FriendChat(msg)
			if err != nil {
				c.write(MsgResp{"error", err.Error()})
			}
		case "group": //群聊
			err := c.GroupChat(msg)
			if err != nil {
				c.write(MsgResp{"error", err.Error()})
			}
		case "updateSign": //修改签名
			sign, ok := msg.Data.(string)
			if !ok {
				fmt.Println("It's not ok for type sign")
			}
			user.Model.Data("sign", sign).Where("id=?", userID).Update()
		case "updateImStatus": //切换状态,在线或者隐身
			imStatus, ok := msg.Data.(string)
			if !ok {
				fmt.Println("It's not ok for type string")
			}
			user.Model.Data("im_status", imStatus).Where("id=?", userID).Update()
			//通知好友
			var msgType string
			if imStatus == "online" {
				msgType = "online"
			} else if imStatus == "hide" {
				msgType = "offline"
			}
			err := c.writeFriends(userID, &MsgResp{
				Type: msgType,
				Data: userID,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

//write 发送消息给当前客户端
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
	userID, e := user_token.GetUserID(data)
	if e != nil {
		return 0, e
	}
	return userID, nil
}

//closeConn 断开连接处理
func (c *Controller) closeConn(ws *ghttp.WebSocket) {
	userID := users.Get(ws)
	fmt.Println(userID)
	fmt.Printf("close user_id:%d\n", userID)
	users.Remove(ws)
	//状态修改为下线，并通知好友
	if userID != 0 {
		userID, ok := userID.(uint)
		if !ok {
			fmt.Println("It's not ok for type uint")
			return
		}
		userIds.Remove(userID)
		user.Model.Data("im_status", "offline").Where("id=?", userID).Update()
		err := c.writeFriends(userID, &MsgResp{
			Type: "offline",
			Data: userID,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

//joinConn 加入连接处理
func (c *Controller) joinConn(userID uint) {
	fmt.Printf("join user_id:%d\n", userID)
	users.Set(c.ws, userID)
	userIds.Set(userID, c.ws)
	//状态修改为在线，并通知好友
	user.Model.Data("im_status", "online").Where("id=?", userID).Update()
	err := c.writeFriends(userID, &MsgResp{
		Type: "online",
		Data: userID,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
