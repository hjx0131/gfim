
import { getToken } from "/resource/js/utils/auth.js";
import { initConfig } from "/resource/js/chat/event.js";

var socket = {}

//创建连接
export function createSocket(url) {
    socket = new WebSocket(url);
    return socket;
}
//发送数据格式
export function socketData(type = "", data = {}) {
    return JSON.stringify({
        type: type,
        data: data,
        token: getToken(),
    })
}
export function sendMsg(type = "", data = {}) {
    socket.send(socketData(type, data))
}
export function wsOpen() {
    sendMsg('confirmJoin', {})
}
export function wsReceive(res) {
    let resp = JSON.parse(res.data);
    //配置im
    if (resp.type === 'initlayim') {
        // 初始化layim
        initConfig()
    }
    //好友聊天
    if (resp.type === 'friend') {
        layui.layim.getMessage(resp.data)
    }
    //获取通知
    if (resp.type == "getNotify") {
        resp.data.forEach(function (val, index, arr) {
            if (val.type === "friend") {
                layui.layim.getMessage(val)
            }
        });
    }
    //群聊
    if (resp.type === 'group') {
        layui.layim.getMessage(resp.data)
    }
    //好友上线
    if (resp.type === 'online') {
        layui.layim.setFriendStatus(resp.data, 'online'); //设置指定好友在线，即头像取消置灰
    }
    //好友离线
    if (resp.type === 'offline') {
        layui.layim.setFriendStatus(resp.data, 'offline'); //设置指定好友在线，即头像取消置灰
    }
    //无效token
    if (resp.type === "invalid_token") {
        layui.layer.msg(resp.data, function () {
            //do something
            window.location.href = "/signIn"

        });
        return;
    }
    //数据统计
    if (resp.type === "count") {
        $("#onlineTotal").html("当前在线人数:" + resp.data.total)
    }
}
export function wsError(event) {
}

export function wsClose(event) {

}
export function socketEvent() {
    socket.onopen = function (event) {
        wsOpen(event);
    }
    socket.onmessage = function (event) {
        wsReceive(event);
    }
    socket.onerror = function (event) {
        wsError(event)
    };
    socket.onclose = function (event) {
        wsClose(event)
    };
}
export function getSocket() {
    return socket;
}