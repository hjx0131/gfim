package user_token

import (
	"errors"
	"gfim/app/model/user_token"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

//GetIDInput 获取ID所需要的参数
type GetIDInput struct {
	Token string `v:"required#token不能为空"`
}

//GetUserID 获取用户ID
func GetUserID(data *GetIDInput) (uint, error) {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return 0, errors.New(e.String())
	}
	one, err := user_token.FindOne("token=?", data.Token)
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
