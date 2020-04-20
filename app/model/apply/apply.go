package apply

//GetListAndTotal 获取列表和总数
func GetListAndTotal(userID uint, page, limit int) ([]*Entity, int, error) {
	list, err := Model.
		Where("to_user_id", userID).
		Or("from_user_id", userID).
		Order("id desc").
		Page(page, limit).
		All()
	if err != nil {
		return nil, 0, err
	}
	count, err := Model.
		Where("to_user_id", userID).
		Or("from_user_id", userID).
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

//GetStateList 获取状态列表
func GetStateList() map[uint]string {
	return map[uint]string{
		1: "待验证",
		2: "已同意",
		3: "已拒绝",
	}
}

//GetStateText 获取状态描述
func (entity *Entity) GetStateText() string {
	list := GetStateList()
	return list[entity.State]
}
