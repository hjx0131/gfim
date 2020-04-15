package user

import (
	"errors"
	"gfim/app/model/friend"
	"gfim/app/model/user"
	"gfim/app/model/user_token"

	"gfim/library/auth"

	"github.com/gogf/gf/crypto/gmd5"

	"gfim/app/service/friend_group"
	"gfim/app/service/group"
	"gfim/library/avatar"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gvalid"
)

const (
	//TokenValidTime token有效时间
	TokenValidTime = 60 * 60 * 24 * 30
)

//SignUpInput 注册输入参数
type SignUpInput struct {
	Username  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Nickname  string `v:"required#昵称不能为空"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
}

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
	//单点登录,将之前的token设置为失效
	user_token.InvalidToken(one.Id)
	//生成token，添加到user_token表
	token := auth.CreateToken()
	now := gtime.Timestamp()
	_, e := user_token.Insert(g.Map{
		"token":       token,
		"user_id":     one.Id,
		"create_time": now,
		"expire_time": now + TokenValidTime,
	})
	if e != nil {
		return "", e
	}
	//更新登录状态
	signInUpdate(one)
	return token, nil
}

//SignUp 用户注册
func SignUp(data *SignUpInput, ip string) error {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.String())
	}
	//检测帐号是否已经存在
	count, err := user.Model.Where("username=?", data.Username).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("用户名已存在")
	}
	salt := createSalt()
	password := encryptPassword(data.Password, salt)
	now := gtime.Timestamp()
	user.Insert(g.Map{
		"username":  data.Username,
		"nickname":  data.Nickname,
		"password":  password,
		"salt":      salt,
		"avatar":    avatar.NewRandom(data.Username),
		"join_time": now,
		"join_ip":   ip,
	})
	return nil
}

//signInUpdate 登录成功后更新状态
func signInUpdate(u *user.Entity) {
	now := gtime.Timestamp()
	user.Update(g.Map{
		"login_time":  now,
		"prev_time":   u.LoginTime,
		"update_time": now,
	})
}

//checkPassword  检验密码，
func checkPassword(pwd, checkpwd, salt string) bool {
	final := encryptPassword(checkpwd, salt)
	return final == pwd
}

//createSalt 生成随机密码盐
func createSalt() string {
	return grand.Letters(6)
}

//encryptPassword 加密密码 通过md5进行第一次加密后的值拼接密码盐后，再进行第二次md5加密
func encryptPassword(pwd, salt string) string {
	first, _ := gmd5.EncryptString(pwd)
	final, _ := gmd5.EncryptString(first + salt)
	return final
}

//Mine 主面板中的个人信息
type Mine struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
}

//Init 主面板
type Init struct {
	Mine   *Mine                     `json:"mine"`
	Friend []*friend_group.GroupInfo `json:"friend"`
	Group  []*group.Info             `json:"group"`
}

//Profile 主面板
func Profile(ID uint) (*Init, error) {
	mine, e := GetMine(ID)
	if e != nil {
		return nil, e
	}
	flist, e := friend_group.GetListByUserID(ID)
	if e != nil {
		return nil, e
	}
	glist, e := group.GetListByUserID(ID)
	if e != nil {
		return nil, e
	}
	data := &Init{
		Mine:   mine,
		Friend: flist,
		Group:  glist,
	}
	return data, nil
}

//GetMine 获取主面板个人信息
func GetMine(ID uint) (*Mine, error) {
	u, e := user.FindOne("id=?", ID)
	if e != nil {
		return nil, e
	}
	if u == nil {
		return nil, errors.New("未找到该用户")
	}
	mine := &Mine{
		Username: u.Nickname,
		ID:       u.Id,
		Status:   u.ImStatus,
		Sign:     u.Sign,
		Avatar:   u.Avatar,
	}
	return mine, nil
}

//SearchRequst 查找用户所需要的参数
type SearchRequst struct {
	Wd    string `json:"wd"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

//SearchResp 查找用户返回格式
type SearchResp struct {
	List  []*Mine `json:"list"`  //列表
	Count int     `json:"count"` //总数
}

//Search 查找用户
func Search(req *SearchRequst, userID uint) (*SearchResp, error) {
	where := make(map[interface{}]interface{})
	if req.Wd != "" {
		where["nickname like ?"] = "%" + req.Wd + "%"
	}
	list, err := SearchList(where, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	count, err := SearchCount(where)
	if err != nil {
		return nil, err
	}
	resp := &SearchResp{
		List:  list,
		Count: count,
	}
	return resp, nil
}

//SearchList 查找用户列表
func SearchList(where map[interface{}]interface{}, page, limit int) ([]*Mine, error) {
	list, err := user.GetListByWhere(where, page, limit)
	if err != nil {
		return nil, err
	}
	res := make([]*Mine, len(list))
	if list != nil {
		for index, item := range list {
			res[index] = &Mine{
				ID:       item.Id,
				Username: item.Nickname,
				Status:   item.ImStatus,
				Sign:     item.Sign,
				Avatar:   item.Avatar,
			}
		}
	}
	return res, nil
}

//SearchCount 查找用户总数
func SearchCount(where map[interface{}]interface{}) (int, error) {
	count, err := user.GetCountByWhere(where)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//Recommend  推荐用户
func Recommend(req *SearchRequst, userID uint) (*SearchResp, error) {
	where := make(map[interface{}]interface{})
	friends, _ := friend.GetFriendUserIds(userID)
	ids := make([]uint, len(friends))
	if friends != nil {
		for index, item := range friends {
			ids[index] = item.Uint()
		}
	}
	ids = append(ids, userID)
	//自己和好友不推荐
	where["id not in (?)"] = ids
	if req.Wd != "" {
		where["nickname like ?"] = "%" + req.Wd + "%"
	}
	list, err := SearchList(where, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	count, err := SearchCount(where)
	if err != nil {
		return nil, err
	}
	resp := &SearchResp{
		List:  list,
		Count: count,
	}
	return resp, nil
}
