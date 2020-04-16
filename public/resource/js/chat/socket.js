
import { getToken } from "../utils/auth.js";

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
export {
    socket,
}