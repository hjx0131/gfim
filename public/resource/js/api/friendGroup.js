import request from "/resource/js/utils/request.js";

/**
 * 新增
 * @param {} params 
 */
export function save(params) {
    return request({
        url: '/api/friendGroup/save',
        method: 'post',
        params
    })
}