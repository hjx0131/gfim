package user

//GetListByWhere 根据条件获取列表
func GetListByWhere(where map[interface{}]interface{}, page, limit int) ([]*Entity, error) {
	list, e := Model.
		Where(where).
		Page(page, limit).
		FindAll()
	if e != nil {
		return nil, e
	}
	return list, e
}

//GetCountByWhere 根据条件获取总数
func GetCountByWhere(where map[interface{}]interface{}) (int, error) {
	count, err := Model.
		Where(where).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
