package systemConfigBLL

import (
	"errors"
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/dal/worldDAL/modelsDAL/systemConfigDAL"
	"github.com/Jordanzuo/GameServer_Go/src/initError"
	"github.com/Jordanzuo/GameServer_Go/src/model/text/systemConfigKey"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/models/systemConfig"
	"strconv"
	"time"
)

var (
	// 系统配置项列表
	systemConfigList = make(map[string]*systemConfig.SystemConfigItem)

	// 设备统计地址
	DeviceStatisticsUrl string

	// 邮件保存的天数
	EmailRetainDays int

	// 战斗服务器地址
	FightServerUrl string

	// 好友消息有效天数
	FriendMessageValidDays int

	// 好友赠送耐力有效天数
	GrantStaValidDays int

	// 登陆服务器地址
	LoginServerUrl string

	// 是否开启邮件发送功能
	MailSendIfOpen bool

	// 邮件接收者
	MailSendTo string

	// 邮件发送Host
	MailHost string

	// 邮件发送账号
	MailAddress string

	// 邮件发送密码
	MailPassword string

	// 队列中允许的最大的消息数量，超过这个数量就会写到文件中
	MessageMaxCount int

	// 队列中允许的最小的消息数量，低于这个数量就会将文件中的信息读入到内存中
	MessageMinCount int

	// 处理消息的线程数量
	MessageThreadCount int

	// 玩家过期所需的小时数，单位：小时
	PlayerExpireNeedHour int

	// 积分墙推广地址
	PromotionActivateUrl string

	// 积分墙检测设备是否激活的接口
	PromotionCheckUrl string

	// 请求日志的最大保存量
	RequestLogMaxCount int

	// 请求时长的临界秒数
	RequestThresholdSeconds int

	// 奖励保留天数
	RewardRetainDays int

	// 事务等级（0：不开启事务；1：开启资源事务；2：开启全部事务）
	TransactionLevel int

	// 活动中心地址
	ActivityCenter string

	// 工会战地址
	GuildWarUrl string

	// Charge API允许访问的IP
	AllowIps string

	// 激活码验证地址
	ActiveCodeRequestAddress string

	// 激活码验证Key
	ActiveCodeRequestSignKey string

	// 游戏代码
	GameCode string

	// 奖励中心批量奖励钻石的最大数量
	RewardAPMaxDiamond int

	// 奖励中心个人奖励钻石的最大数量
	RewardOPMaxDiamond int

	// 充值订单接口
	ChargeUrl string

	// 玩家Session过期时间（min）
	SessionExpireMinutes int

	// 错误日志发送大小
	SendErrorLogSize int

	// 错误日志发送天数
	SendErrorLogDays int

	// 是否存储组队副本队伍信息
	IsMemoryTeamcopyTeams bool

	// 保存服务器列表API
	SaveUserServerHistory string

	// 检查客户端过期的时间间隔，单位（秒）
	CheckExpiredInterval time.Duration

	// 客户端过期的秒数
	ClientExpiredSeconds time.Duration
)

func init() {
	rows, err := systemConfigDAL.GetList()
	if err != nil {
		initError.RegisterInitError(err)
		return
	}

	if rows == nil {
		initError.RegisterInitError(errors.New("未找到SystemConfig的配置项"))
		return
	}

	for rows.Next() {
		var configKey string
		var configValue string
		err = rows.Scan(&configKey, &configValue)
		if err != nil {
			initError.RegisterInitError(err)
			continue
		}

		// 构造对象并添加到列表中
		systemConfigItem := systemConfig.New(configKey, configValue)
		systemConfigList[systemConfigItem.ConfigKey] = systemConfigItem
	}

	// 检测数据
	check()
}

// for test
func init() {
	CheckExpiredInterval = time.Duration(60)
	ClientExpiredSeconds = time.Duration(60)
}

func check() {
	var value *systemConfig.SystemConfigItem
	var ok = false

	if value, ok = systemConfigList[systemConfigKey.CheckExpiredInterval]; !ok {
		initError.RegisterInitError(errors.New(fmt.Sprintf("没有找到%s的配置", systemConfigKey.CheckExpiredInterval)))
	} else {
		if checkExpiredInterval_int, err := strconv.Atoi(value.ConfigValue); err != nil {
			initError.RegisterInitError(err)
		} else {
			CheckExpiredInterval = time.Duration(checkExpiredInterval_int)
		}
	}

	if value, ok = systemConfigList[systemConfigKey.ClientExpiredSeconds]; !ok {
		initError.RegisterInitError(errors.New(fmt.Sprintf("没有找到%s的配置", systemConfigKey.ClientExpiredSeconds)))
	} else {
		if clientExpiredSeconds_int, err := strconv.Atoi(value.ConfigValue); err != nil {
			initError.RegisterInitError(err)
		} else {
			ClientExpiredSeconds = time.Duration(clientExpiredSeconds_int)
		}
	}
}
