/*
服务器组配置包
*/
package serverGroupSetting

// 服务器组设置配置对象
type ServerGroupSettingConfig struct {
	// 服务器组Id
	serverGroupId int

	// ManageCenter激活服务器地址
	manageCenterUrl string
}

// 创建新的服务器组设置配置对象
func New(serverGroupId int, manageCenterUrl string) *ServerGroupSettingConfig {
	return &ServerGroupSettingConfig{
		serverGroupId:   serverGroupId,
		manageCenterUrl: manageCenterUrl,
	}
}

// 获取服务器组Id
// 返回值：
// 服务器组Id
func (sgsc *ServerGroupSettingConfig) ServerGroupId() int {
	return sgsc.serverGroupId
}

// 获取ManageCenter激活服务器地址
// 返回值：
// ManageCenter激活服务器地址
func (sgsc *ServerGroupSettingConfig) ManageCenterUrl() string {
	return sgsc.manageCenterUrl
}
