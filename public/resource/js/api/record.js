import request from "/resource/js/utils/request.js";

/**
 * 获取聊天记录
 * @param {} params 
 */
export function getRecord(params) {
    return request({
        url: '/api/record/getData',
        method: 'get',
        params
    })
}