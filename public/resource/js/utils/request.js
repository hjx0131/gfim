import { getToken } from "./auth.js";

// 创建axios实例
const service = axios.create({
    baseURL: '/', // api 的 base_url
    timeout: 0 // 请求超时时间
})

// request拦截器
service.interceptors.request.use(
    config => {
        config.headers['X-Auth-Token'] = getToken() // 让每个请求携带自定义token 请根据实际情况自行修改
        return config
    },
    error => {
        // Do something with request error
        console.log(error) // for debug
        Promise.reject(error)
    }
)
layui.use('layer', function () {
    var layer = layui.layer;
    // response 拦截器
    service.interceptors.response.use(
        response => {
            /**
             * code为非0是抛错 可结合自己业务进行修改
             */
            console.log(Promise, 'promise')
            const res = response.data
            if (res.code !== 0) {
                layer.msg(res.msg, { time: 3000 })
                // 50008:非法的token; 50012:其他客户端登录了;  50014:Token 过期了;
                if (res.code === 2) {
                    layer.confirm('token无效，是否重新登录', function (index) {
                        layer.close(index);
                        window.location.href = '/signIn';
                    });
                }
                return Promise.reject('error')
            } else {
                return response.data
            }
        },
        error => {
            console.log('err' + error) // for debug
            layer.msg(res.msg, { time: 3000 })
            return Promise.reject(error)
        }
    )

});

export default service