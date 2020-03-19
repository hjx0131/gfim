package user_token

import (
	"errors"
	"gfim/app/model/user_token"

	"github.com/gogf/gf/os/gtime"
)

//GetUserID 获取用户ID
func GetUserID(token string) (uint, error) {
	one, err := user_token.FindOne("token=?", token)
	if err != nil {
		return 0, err
	}
	if one == nil {
		return 0, errors.New("token无效")
	}
	now := gtime.Timestamp()
	if now > int64(one.Expiretime) {
		return 0, errors.New("token已过期")
	}
	return one.UserId, nil
}
