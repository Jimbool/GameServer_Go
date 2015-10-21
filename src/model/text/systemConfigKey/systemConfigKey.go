package systemConfigKey

const (
	// 设备统计地址
	DeviceStatisticsUrl = "DeviceStatisticsUrl"

	// 邮件保存的天数
	EmailRetainDays = "EmailRetainDays"

	// 战斗服务器地址
	FightServerUrl = "FightServerUrl"

	// 好友消息有效天数
	FriendMessageValidDays = "FriendMessageValidDays"

	// 好友赠送耐力有效天数
	GrantStaValidDays = "GrantStaValidDays"

	// 登陆服务器地址
	LoginServerUrl = "LoginServerUrl"

	// 是否开启邮件发送功能
	MailSendIfOpen = "MailSendIfOpen"

	// 邮件接收者
	MailSendTo = "MailSendTo"

	// 邮件发送Host
	MailHost = "MailHost"

	// 邮件发送账号
	MailAddress = "MailAddress"

	// 邮件发送密码
	MailPassword = "MailPassword"

	// 队列中允许的最大的消息数量，超过这个数量就会写到文件中
	MessageMaxCount = "MessageMaxCount"

	// 队列中允许的最小的消息数量，低于这个数量就会将文件中的信息读入到内存中
	MessageMinCount = "MessageMinCount"

	// 处理消息的线程数量
	MessageThreadCount = "MessageThreadCount"

	// 玩家过期所需的小时数，单位：小时
	PlayerExpireNeedHour = "PlayerExpireNeedHour"

	// 积分墙推广地址
	PromotionActivateUrl = "PromotionActivateUrl"

	// 积分墙检测设备是否激活的接口
	PromotionCheckUrl = "PromotionCheckUrl"

	// 请求日志的最大保存量
	RequestLogMaxCount = "RequestLogMaxCount"

	// 请求时长的临界秒数
	RequestThresholdSeconds = "RequestThresholdSeconds"

	// 奖励保留天数
	RewardRetainDays = "RewardRetainDays"

	// 事务等级（0：不开启事务；1：开启资源事务；2：开启全部事务）
	TransactionLevel = "TransactionLevel"

	// 活动中心地址
	ActivityCenter = "ActivityCenter"

	// 工会战地址
	GuildWarUrl = "GuildWarUrl"

	// Charge API允许访问的IP
	AllowIps = "AllowIps"

	// 激活码验证地址
	ActiveCodeRequestAddress = "ActiveCodeRequestAddress"

	// 激活码验证Key
	ActiveCodeRequestSignKey = "ActiveCodeRequestSignKey"

	// 游戏代码
	GameCode = "GameCode"

	// 奖励中心批量奖励钻石的最大数量
	RewardAPMaxDiamond = "RewardAPMaxDiamond"

	// 奖励中心个人奖励钻石的最大数量
	RewardOPMaxDiamond = "RewardOPMaxDiamond"

	// 充值订单接口
	ChargeUrl = "ChargeUrl"

	// 玩家Session过期时间（min）
	SessionExpireMinutes = "SessionExpireMinutes"

	// 错误日志发送大小
	SendErrorLogSize = "SendErrorLogSize"

	// 错误日志发送天数
	SendErrorLogDays = "SendErrorLogDays"

	// 是否存储组队副本队伍信息
	IsMemoryTeamcopyTeams = "IsMemoryTeamcopyTeams"

	// 保存服务器列表API
	SaveUserServerHistory = "SaveUserServerHistory"

	// 检查客户端过期的时间间隔，单位（秒）
	CheckExpiredInterval = "CheckExpiredInterval"

	// 客户端过期的秒数
	ClientExpiredSeconds = "ClientExpiredSeconds"
)
