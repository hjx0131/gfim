package user

//GetListByWhere 关键字搜索用户
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
