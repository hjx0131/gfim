import { getToken } from "/resource/js/utils/auth.js";
import { redirect } from "/resource/js/utils/tools.js";

//白名单
const whiteList = ['/signIn', '/signUp']
export function check() {
    if (getToken()) {
        return true
    }
    var path = window.location.pathname
    if (whiteList.indexOf(path) !== -1) {
        return true
    }
    redirect("/signIn")
}