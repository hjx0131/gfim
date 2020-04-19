
import { getToken } from "/resource/js/utils/auth.js";
import { initConfig } from "/resource/js/chat/event.js";
import { redirect } from "/resource/js/utils/tools.js";

var socket = {}

//创建连接
export function createSocket(url) {
    socket = new WebSocket(url);
    window.ws = socket
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
    if (resp.error === true) {
        layui.layer.msg(resp.data)
        return;
    }
    //成功处理返回提示信息
    if (resp.type === 'success') {
        layui.layer.msg(resp.data)
    }
    if (resp.type === "applyCount") {
        console.log(layui.layim)
        layui.layim.msgbox(resp.data)
    }
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
        layui.layer.confirm(resp.data + ',是否重新登录?', function (index) {
            //do something
            redirect("signIn")
            layer.close(index);
        });
        return;
    }
    //数据统计
    if (resp.type === "count") {
        $("#onlineTotal").html("当前在线人数:" + resp.data.total)
    }
}
export function wsError(event) {
    console.log(event)
    layui.layer.msg("连接出错")

}

export function wsClose(event) {
    console.log(event)
    // layer.confirm('连接已断开', function (index) {
    //     //do something
    //     layer.close(index);
    // });

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