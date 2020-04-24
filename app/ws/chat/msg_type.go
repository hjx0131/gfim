package chat

const (
	//SystemError 系统错误
	SystemError = "system_error"

	//Success 成功处理
	Success = "success"

	//InvalidParamterFormat 无效参数格式
	InvalidParamterFormat = "invalid_parameter_format"

	//ParameterValidationFailed  参数校验失败
	ParameterValidationFailed = "parameter_validation_failed "

	//InvalidToken 无效token
	InvalidToken = "invalid_token"

	//Ping 心跳检测
	Ping = "ping"

	//Pong 心跳检测
	Pong = "pong"

	//ConfirmJoin 确认连接
	ConfirmJoin = "confirm_join"

	//InitLayimConfig 初始化layim配置
	InitLayimConfig = "init_layim_config"

	//NotifyRecord 聊天消息通知
	NotifyRecord = "notify_record"

	//NoReadApply 未读好友申请
	NoReadApply = "no_read_apply"

	//FriendChat 好友聊天
	FriendChat = "friend"

	//GroupChat 群聊天
	GroupChat = "group"

	//UpdateSign 修改签名
	UpdateSign = "update_sign"

	//UpdateStatus 修改在线状态
	UpdateStatus = "update_status"

	//Online 在线
	Online = "online"

	//Offline 离线
	Offline = "offline"

	//Hide  隐身
	Hide = "hide"

	//ApplyFriend 好友申请
	ApplyFriend = "apply_friend"

	//AgreeFriend 同意好友申请
	AgreeFriend = "agree_friend"

	//RefuseFriend 拒绝好友申请
	RefuseFriend = "refuse_friend"

	//CountData 数据统计
	CountData = "count_data"

	//AppendFriend 追加好友
	AppendFriend = "append_friend"

	//ApplyGroup 群组申请
	ApplyGroup = "apply_group"

	//AgreeGroup 同意好友申请
	AgreeGroup = "agree_group"

	//RefuseGroup 拒绝好友申请
	RefuseGroup = "refuse_group"

	//AppendGroup 追加群组
	AppendGroup = "append_group"
)
