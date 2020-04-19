import request from "/resource/js/utils/request.js";

/**
 * 登录
 * @param {} params 
 */
export function signIn(params) {
    return request({
        url: '/api/user/signIn',
        method: 'post',
        params
    })
}
/**
 * 注册
 * @param {} params 
 */
export function signUp(params) {
    return request({
        url: '/api/user/signUp',
        method: 'post',
        params
    })
}
/**
 * 注销
 * @param {} params 
 */
export function logout(params) {
    return request({
        url: '/api/user/logout',
        method: 'post',
        params
    })
}