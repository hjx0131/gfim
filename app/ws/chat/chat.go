package chat

import (
	"fmt"
	"gfim/app/model/user"
	"gfim/app/service/user_token"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
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
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
	Error bool        `json:"error"`
}

//Controller 控制器
type Controller struct {
}

//CountDataResp 统计数据返回格式
type CountDataResp struct {
	Total int `json:"total"`
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
	for {
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			//断开连接处理
			c.closeConn(ws)
			getCountData()
			break
		}
		// json解析
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			writeByWs(ws, &MsgResp{
				Type:  InvalidParamterFormat,
				Data:  "消息格式不正确: " + err.Error(),
				Error: true,
			})
			continue
		}
		// 数据校验
		if err := gvalid.CheckStruct(msg, nil); err != nil {
			writeByWs(ws, &MsgResp{
				Type:  ParameterValidationFailed,
				Data:  err.String(),
				Error: true,
			})
			continue
		}
		//频繁心跳检测，不做token验证
		if msg.Type == Ping {
			writeByWs(ws, &MsgResp{
				Type: Pong,
			})
			continue
		}
		// 检验token
		userID, err := getUserID(msg.Token)
		if err != nil {
			writeByWs(ws, &MsgResp{
				Type:  InvalidToken,
				Data:  err.Error(),
				Error: true,
			})
			//ws.Close()
			continue
		}
		//发送消息
		switch msg.Type {

		case ConfirmJoin: //确认连接
			c.joinConn(ws, userID)
			writeByWs(ws, &MsgResp{
				Type: InitLayimConfig,
			})
			//统计数据
			getCountData()

		case NotifyRecord: //获取通知
			c.notifyUserRecord(userID)

		case NoReadApply: //获取未读验证消息
			c.NoReadApplyCount(userID)

		case Ping: //心跳检测
			writeByWs(ws, &MsgResp{
				Type: Pong,
			})

		case FriendChat: //好友聊天
			err := c.FriendChat(msg)
			if err != nil {
				writeByWs(ws, &MsgResp{
					Type:  SystemError,
					Data:  err.Error(),
					Error: true,
				})
			}

		case GroupChat: //群聊
			err := c.GroupChat(msg)
			if err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			}

		case UpdateSign: //修改签名
			sign, ok := msg.Data.(string)
			if !ok {
				fmt.Println("It's not ok for type sign")
			}
			user.Model.Data("sign", sign).Where("id=?", userID).Update()
			writeByWs(ws, &MsgResp{Success, "签名修改成功", false})

		case UpdateStatus: //切换状态,在线或者隐身
			imStatus, ok := msg.Data.(string)
			if !ok {
				fmt.Println("It's not ok for type string")
			}
			user.Model.Data("im_status", imStatus).Where("id=?", userID).Update()
			//通知好友
			msgType := ""
			if imStatus == Online {
				msgType = Online
			} else if imStatus == Hide {
				msgType = Offline
			}
			err := c.writeFriends(userID, &MsgResp{
				Type: msgType,
				Data: userID,
			})
			if err != nil {
				fmt.Println(err.Error())
			}

		case ApplyFriend: //好友申请
			if err := c.apply(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "申请成功", false})
			}

		case AgreeFriend: //同意好友申请
			if err := c.agree(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "操作成功", false})
			}

		case RefuseFriend: //拒绝好友申请
			if err := c.refuse(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "操作成功", false})
			}
		case ApplyGroup: //群组申请
			if err := c.applyGroup(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "申请成功", false})
			}

		case AgreeGroup: //同意入群
			if err := c.agree(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "操作成功", false})
			}

		case RefuseGroup: //拒绝入群
			if err := c.refuse(userID, msg); err != nil {
				writeByWs(ws, &MsgResp{SystemError, err.Error(), true})
			} else {
				writeByWs(ws, &MsgResp{Success, "操作成功", false})
			}
		}
	}
}

//getCountData 获取统计数据
func getCountData() {
	writeAll(&MsgResp{
		Type: CountData,
		Data: &CountDataResp{
			Total: users.Size(),
		},
	})
}

//writeByWs 指定ws发送消息
func writeByWs(ws *ghttp.WebSocket, msg *MsgResp) error {
	data, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	if ws != nil {
		ws.WriteMessage(ghttp.WS_MSG_TEXT, data)
	}
	return nil
}

//writeAll 发送给所有客户端
func writeAll(msg *MsgResp) {
	if users != nil {
		// 遍历map
		users.Iterator(func(k interface{}, v interface{}) bool {
			writeByWs(k.(*ghttp.WebSocket), msg)
			return true
		})
	}
}

func (c *Controller) writeByUserID(userID uint, msg *MsgResp) {
	ws := userIds.Get(userID)
	if ws != nil {
		writeByWs(ws.(*ghttp.WebSocket), msg)
	}
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
	users.Remove(ws)
	//状态修改为下线，并通知好友
	if userID != 0 {
		userID, ok := userID.(uint)
		if !ok {
			fmt.Println("It's not ok for type uint")
			return
		}
		if ws == userIds.Get(userID) {
			userIds.Remove(userID)
			user.Model.Data("im_status", Offline).Where("id=?", userID).Update()
			err := c.writeFriends(userID, &MsgResp{
				Type: Offline,
				Data: userID,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

//joinConn 加入连接处理
func (c *Controller) joinConn(ws *ghttp.WebSocket, userID uint) {
	fmt.Printf("join user_id:%d\n", userID)
	users.Set(ws, userID)
	//如果该帐号已在其他客户端登录，通知客户端下线
	oldWs := getWsByUserID(userID)
	if oldWs != nil {
		writeByWs(oldWs, &MsgResp{
			Type: InvalidToken,
			Data: "该帐号已在其他设备登录",
		})
		oldWs.Close()
	}
	userIds.Set(userID, ws)
	//状态修改为在线，并通知好友
	user.Model.Data("im_status", Online).Where("id=?", userID).Update()
	err := c.writeFriends(userID, &MsgResp{
		Type: Online,
		Data: userID,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//getWsByUserID 根据用户id获取ws连接
func getWsByUserID(userID uint) *ghttp.WebSocket {
	ws, ok := userIds.Get(userID).(*ghttp.WebSocket)
	if ok {
		return ws
	}
	return nil
}
