package friend_group

import (
	"errors"
	"gfim/app/model/friend"
	"gfim/app/model/friend_group"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

//GroupInfo 好友群组结构体,前台需要的格式
type GroupInfo struct {
	ID        uint          `json:"id"`
	GroupName string        `json:"groupname"`
	List      []*FriendInfo `json:"list"`
}

//FriendInfo 好友结构体，前台需要的格式
type FriendInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Status   string `json:"status"`
}

// GetListByUserID 根据用户ID获取好友分组列表
func GetListByUserID(userID uint) ([]*GroupInfo, error) {
	//好友群组列表
	glist, err := friend_group.GetListByUserID(userID)
	if err != nil {
		return nil, err
	}
	//好友列表
	flist, _ := friend.GetListByUserID(userID)
	res := make([]*GroupInfo, len(glist))
	if glist != nil {
		for index, item := range glist {
			f := make([]*FriendInfo, 0)
			if flist != nil {
				for _, val := range flist {
					if val["friend_group_id"].Uint() == item.Id {
						imStatus := val["im_status"].String()
						//隐身状态显示离线
						if imStatus == "hide" {
							imStatus = "offline"
						}
						f = append(f, &FriendInfo{
							ID:       val["friend_id"].Uint(),
							Username: val["nickname"].String(),
							Avatar:   val["avatar"].String(),
							Sign:     val["sign"].String(),
							Status:   imStatus,
						})
					}
				}
			}
			res[index] = &GroupInfo{
				ID:        item.Id,
				GroupName: item.Name,
				List:      f,
			}
		}
	}
	return res, nil
}

//SaveReq 保存请求参数
type SaveReq struct {
	UserID uint   `json:"user_id" v:"user_id@required#创建人不能为空"`
	Name   string `json:"name" v:"name@required#分组名不能为空"`
}

//Save 保存数据
func Save(req *SaveReq) error {
	// 数据校验
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return err
	}
	count, err := friend_group.Model.
		Where("user_id in(?) ", g.Slice{0, req.UserID}).
		Where("name", req.Name).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("分组名已存在")
	}
	now := gtime.Timestamp()
	_, err = friend_group.Model.
		Data(g.Map{
			"user_id":     req.UserID,
			"name":        req.Name,
			"create_time": now,
		}).
		Insert()
	if err != nil {
		return err
	}
	return nil
}
