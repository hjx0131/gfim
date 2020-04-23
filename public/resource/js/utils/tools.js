export function getParams(key = '') {
    var result = {};
    var paramStr = encodeURI(window.document.location.search);
    if (paramStr) {
        var params = paramStr.substring(1).split('&');
        params.map(v => result[v.split("=")[0]] = v.split("=")[1])
    }
    if (key) {
        return result[key] ? result[key] : ''
    }
    return result
}
export function redirect(path) {
    window.location.href = path;
}

export function getwsURL() {
    return $("#wsURL").val()
}