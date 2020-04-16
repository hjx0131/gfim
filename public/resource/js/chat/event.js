import { sendMsg } from "/resource/js/chat/socket.js";
import { getToken } from "../utils/auth.js";

export function initConfig() {
    layui.use('layim', function (layim) {
        // 初始化layim
        layim.config({
            init: {
                url: "/api/user/profile",
                type: 'get', //默认get，一般可不填
                data: {
                    token: getToken()//额外参数
                },
            },
            //获取群员接口（返回的数据格式见下文）
            members: {
                url: '/api/group/userList', //接口地址（返回的数据格式见下文）
                type: 'get', //默认get，一般可不填
                data: {
                    token: getToken()
                }, //额外参数
            },
            //上传图片接口（返回的数据格式见下文），若不开启图片上传，剔除该项即可
            uploadImage: {
                url: '',//接口地址
                type: 'post',//默认post
            },
            //上传文件接口（返回的数据格式见下文），若不开启文件上传，剔除该项即可
            uploadFile: {
                url: '', //接口地址
                type: 'post', //默认post
            },
            //扩展工具栏，下文会做进一步介绍（如果无需扩展，剔除该项即可）
            tool: [{
                alias: 'code', //工具别名
                title: '代码', //工具名称
                icon: '&#xe64e;', //工具图标，参考图标文档
            }],
            msgbox: '/msgbox', //消息盒子页面地址，若不开启，剔除该项即可
            find: '/find', //发现页面地址，若不开启，剔除该项即可
            chatLog: '/chatlog', //聊天记录页面地址，若不开启，剔除该项即可
        });
    })
}
/**
 * im初始化完成
 */
export function ready() {
    layui.layim.on('ready', function (options) {
        let html = layui.laytpl(demo.innerHTML).render({
            data: layui.layim.cache().mine
        });
        $('#view').html(html);
        //获取未通知的好友消息
        sendMsg('getNotify', {})
    });
}
//修改签名
export function updateSign() {
    //监听修改签名
    layui.layim.on('sign', function (value) {
        sendMsg('updateSign', value)
    });
}

//修改在线状态
export function updateImStatus() {
    layui.layim.on('online', function (status) {
        sendMsg('updateImStatus', status)
    });
}
//发送消息
export function sendMessage() {
    //监听发送消息
    layui.layim.on('sendMessage', function (res) {
        if (res.to.type == "friend") {
            //好友消息
            sendMsg('friend', {
                from_user_id: res.mine.id,
                to_user_id: res.to.id,
                content: res.mine.content
            })
        }
        if (res.to.type == "group") {
            //群消息
            sendMsg('group', {
                user_id: res.mine.id,
                group_id: res.to.id,
                content: res.mine.content
            })
        }
    })
}