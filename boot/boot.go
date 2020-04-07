package boot

import (
	"gfim/app/model/user"
)

func init() {
	user.Model.Data("im_status", "offline").Update()

}
