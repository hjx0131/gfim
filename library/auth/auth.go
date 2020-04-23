package auth

import (
	"strconv"
	"strings"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
)

//KeyPre 缓存键名前缀
const (
	KeyPre = "token_key_pre"
)

//CreateToken 生成token
func CreateToken() string {
	data := strings.ToUpper(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	token, _ := gmd5.EncryptString(data)
	return token
}

//GetToken 获取token
func GetToken(r *ghttp.Request) string {
	token := r.Header.Get("X-Auth-Token")
	if token == "" {
		if queryToken := r.Get("token"); queryToken != nil {
			token = queryToken.(string)
		}
	}
	return token
}
