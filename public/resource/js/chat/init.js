import { createSocket, socketEvent } from "/resource/js/chat/socket.js";
import { ready, updateSign, updateImStatus, sendMessage } from "/resource/js/chat/event.js";
import { getwsURL } from "/resource/js/utils/tools.js";

export function init() {
    layui.use(['layim', 'laytpl', 'layer', 'laytpl'], function () {
        var websocket = createSocket(getwsURL())
        socketEvent(websocket)
        ready()
        updateSign()
        updateImStatus()
        sendMessage()
    })
}
