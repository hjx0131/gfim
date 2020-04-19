package boot

import (
	"gfim/app/model/user"
)

func init() {
	user.Model.
		Where("im_status <>", "offline").
		Data("im_status", "offline").
		Update()
}
