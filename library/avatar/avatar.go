package avatar

import (
	"bytes"

	"github.com/gogf/gf/crypto/gmd5"
)

//NewRandom 生成随机头像
func NewRandom(str string) string {
	//随机头像
	hash, _ := gmd5.EncryptString(str)
	var avatar bytes.Buffer
	avatar.WriteString("https://secure.gravatar.com/avatar/")
	avatar.WriteString(hash)
	avatar.WriteString("?s=128&d=identicon&r=PG")
	return avatar.String()
}
