package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gfim/app/model/user"
	"gfim/app/model/user_token"
	"gfim/library/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

const (
	//TokenValidTime token有效时间
	TokenValidTime = 60 * 60 * 24 * 30
)

//SignIn 用户登录，成功返回token，否则返回nil
func SignIn(username, password string) (string, error) {
	one, err := user.FindOne("username=?", username)
	if err != nil {
		return "", err
	}
	//检验用户是否存在
	if one == nil {
		return "", errors.New("用户不存在")
	}
	//验证密码是否正确
	if checkPassword(one.Password, password, one.Salt) == false {
		return "", errors.New("密码不正确")
	}
	//生成token，添加到user_token表
	token := auth.CreateToken()
	now := gtime.Timestamp()
	_, e := user_token.Insert(g.Map{
		"token":      token,
		"user_id":    one.Id,
		"createtime": now,
		"expiretime": now + TokenValidTime,
	})
	if e != nil {
		return "", e
	}
	//更新登录状态
	signInUpdate(one)
	return token, nil
}

//signInUpdate 登录成功后更新状态
func signInUpdate(u *user.Entity) {
	now := gtime.Timestamp()
	user.Update(g.Map{
		"logintime":  now,
		"prevtime":   u.Logintime,
		"updatetime": now,
	})
}

//checkPassword  检验密码，
func checkPassword(pwd, checkpwd, salt string) bool {
	final := encryptPassword(checkpwd, salt)
	return final == pwd
}

//encryptPassword 加密密码 通过md5进行第一次加密后的值拼接密码盐后，再进行第二次md5加密
func encryptPassword(pwd, salt string) string {
	h := md5.New()
	h2 := md5.New()
	h.Write([]byte(pwd))
	first := hex.EncodeToString(h.Sum(nil))
	h2.Write([]byte(first + salt))
	return hex.EncodeToString(h2.Sum(nil))
}
