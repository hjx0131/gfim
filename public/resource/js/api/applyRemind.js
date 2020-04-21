import request from "/resource/js/utils/request.js";

/**
 * 标记为已读
 * @param {} params 
 */
export function setIsRead(params) {
    return request({
        url: '/api/applyRemind/setIsRead',
        method: 'post',
        params
    })
}
