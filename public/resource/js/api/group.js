import request from "/resource/js/utils/request.js";

/**
 * 获取列表
 * @param {} params 
 */
export function search(params) {
    return request({
        url: '/api/group/search',
        method: 'get',
        params
    })
}

/**
 * 新增
 * @param {} params 
 */
export function save(params) {
    return request({
        url: '/api/group/save',
        method: 'post',
        params
    })
}