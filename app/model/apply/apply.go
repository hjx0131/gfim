package apply

import (
	"github.com/gogf/gf/frame/g"
)

//GetListAndTotal 获取列表和总数
func GetListAndTotal(userID uint, page, limit int) ([]*Entity, int, error) {
	list, err := Model.
		Where(g.Map{
			"to_user_id": userID,
		}).
		Or(g.Map{
			"from_user_id": userID,
		}).
		Order("id desc").
		Page(page, limit).
		All()
	if err != nil {
		return nil, 0, err
	}
	count, err := Model.
		Where(g.Map{
			"to_user_id": userID,
		}).
		Or(g.Map{
			"from_user_id": userID,
		}).
		Count()
	if err != nil {
		return list, 0, err
	}
	return list, count, nil
}

//GetNoHandleCount 获取待处理总数
func GetNoHandleCount(userID uint) (int, error) {
	count, err := Model.
		Where("to_user_id", userID).
		Where("state", 1).
		Count()
	if err != nil {
		return 0, nil
	}
	return count, nil
}
