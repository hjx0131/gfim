import { createSocket, socketEvent } from "/resource/js/chat/socket.js";
import { ready, updateSign, updateImStatus, sendMessage } from "./event.js";
layui.use(['layim', 'laytpl', 'layer', 'laytpl'], function (layim) {
    //初始ws
    createSocket('ws://localhost:8199/chat')
    //监听ws
    socketEvent()
    ready()
    updateSign()
    updateImStatus()
    sendMessage()
})