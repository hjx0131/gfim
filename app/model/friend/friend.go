package friend

//GetListByUserID 获取好友列表
func GetListByUserID(UserID uint) ([]*Entity, error) {
	list, err := Model.As("f").
		InnerJoin("gf_user u", "u.id=f.friend_id").
		Where("f.user_id=?", UserID).
		Fields("f.*,u.nickname,u.avatar,u.status,u.bio").
		All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
