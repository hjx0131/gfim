import { createSocket, socketEvent, getSocket } from "/resource/js/chat/socket.js";
import { ready, updateSign, updateImStatus, sendMessage } from "/resource/js/chat/event.js";
layui.use(['layim', 'laytpl', 'layer', 'laytpl'], function () {
    createSocket('ws://localhost:8199/chat')
    socketEvent()
    ready()
    updateSign()
    updateImStatus()
    sendMessage()
})