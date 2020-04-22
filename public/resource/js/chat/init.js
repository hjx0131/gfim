import { createSocket, socketEvent } from "/resource/js/chat/socket.js";
import { ready, updateSign, updateImStatus, sendMessage } from "/resource/js/chat/event.js";
import { wsUrl } from "/resource/js/env.js";

export function init() {
    layui.use(['layim', 'laytpl', 'layer', 'laytpl'], function () {
        var websocket = createSocket(wsUrl)
        socketEvent(websocket)
        ready()
        updateSign()
        updateImStatus()
        sendMessage()
    })
}
