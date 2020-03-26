package chat

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

//Msg 消息结构体
type Msg struct {
	Type string      `json:"type" v:"type@required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	From string      `json:"name" v:""`
}

//WebSocket ws
func WebSocket(r *ghttp.Request) {
	// 初始化WebSocket请求
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if err = ws.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
