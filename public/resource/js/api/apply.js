import request from "/resource/js/utils/request.js";

/**
 * 登录
 * @param {} params 
 */
export function getData(params) {
    return request({
        url: '/api/apply/getData',
        method: 'post',
        params
    })
}
/**
 * 同意
 * @param {} params 
 */
export function agree(params) {
    return request({
        url: '/api/apply/agree',
        method: 'post',
        params
    })
}
/**
 * 拒绝
 * @param {} params 
 */
export function refuse(params) {
    return request({
        url: '/api/apply/refuse',
        method: 'post',
        params
    })
}