package apply_remind

import (
	"gfim/app/model/apply_remind"

	"github.com/gogf/gf/os/gtime"
)

//SetIsRead 根据用户ID将提醒标记为已读
func SetIsRead(userID uint) {
	data := map[string]interface{}{
		"is_read":   1,
		"read_time": gtime.Timestamp(),
	}
	apply_remind.
		Model.
		Where("user_id", userID).
		Data(data).
		Update()
}
