package user_token

//InvalidToken 将token设为失效
func InvalidToken(userID uint) {
	Model.Where("user_id=?", userID).
		Where("is_valid", 1).
		Data("is_valid", 0).
		Update()
}
