package chat

import (
	"gfim/app/model/apply"
)

//NoHandleApplyCount 获取未处理验证消息
func (c *Controller) NoHandleApplyCount(userID uint) error {
	count, err := apply.GetNoHandleCount(userID)
	if err != nil {
		return err
	}
	if count > 0 {
		c.writeByUserID(userID, &MsgResp{
			Type: "applyCount",
			Data: count,
		})
	}
	return nil
}
