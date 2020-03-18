package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gfim/app/model/user"
	"gfim/library/auth"
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
	//生产token，放到redis缓存中
	token := auth.CreateToken()
	auth.SetTokenID(token, one.Id)
	return token, nil
}

//checkPassword  检验密码， 通过md5进行第一次加密后的值拼接密码盐后，再进行第二次md5加密
func checkPassword(password, checkpassword, salt string) bool {
	h := md5.New()
	h2 := md5.New()
	h.Write([]byte(checkpassword))
	first := hex.EncodeToString(h.Sum(nil))
	h2.Write([]byte(first + salt))
	final := hex.EncodeToString(h2.Sum(nil))
	return final == password
}
