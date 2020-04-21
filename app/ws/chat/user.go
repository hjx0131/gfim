package chat

import (
	"gfim/app/model/apply_remind"
)

//NoHandleApplyCount 获取未处理验证消息
// func (c *Controller) NoHandleApplyCount(userID uint) error {
// 	count, err := apply.GetNoHandleCount(userID)
// 	if err != nil {
// 		return err
// 	}
// 	if count > 0 {
// 		c.writeByUserID(userID, &MsgResp{
// 			Type: "applyCount",
// 			Data: count,
// 		})
// 	}
// 	return nil
// }

//NoReadApplyCount 获取未处理验证消息
func (c *Controller) NoReadApplyCount(userID uint) error {
	count, err := apply_remind.Model.
		Where("user_id", userID).
		Where("is_read", 0).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		c.writeByUserID(userID, &MsgResp{
			Type: NoReadApply,
			Data: count,
		})
	}
	return nil
}
