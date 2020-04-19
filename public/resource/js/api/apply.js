import request from "/resource/js/utils/request.js";

/**
 * 登录
 * @param {} params 
 */
export function getList(params) {
    return request({
        url: '/api/apply/index',
        method: 'post',
        params
    })
}
