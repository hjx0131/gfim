package auth

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

//KeyPre 缓存键名前缀
const (
	KeyPre = "token_key_pre"
)

//CreateToken create token for login user
func CreateToken() string {
	return strings.ToUpper(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
}

//SetTokenID 将uid存放到token中
func SetTokenID(token string, id uint) {
	key := KeyPre + token
	_, err := g.Redis().Do("SET", key, id)
	if err != nil {
		panic(err)
	}
	r, _ := g.Redis().Do("get", key)

	fmt.Println(gconv.String(r))

}
