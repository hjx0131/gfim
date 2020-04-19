import request from "/resource/js/utils/request.js";

/**
 * 查找用户
 * @param {} params 
 */
export function search(params) {
    return request({
        url: '/api/user/search',
        method: 'post',
        params
    })
}
/**
 * 推荐用户
 * @param {} params 
 */
export function recommend(params) {
    return request({
        url: '/api/user/recommend',
        method: 'post',
        params
    })
}
