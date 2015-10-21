package systemConfig

// 系统配置项
type SystemConfigItem struct {
	// 配置的Key
	ConfigKey string

	// 配置的值
	ConfigValue string
}

func New(configKey, configValue string) *SystemConfigItem {
	return &SystemConfigItem{
		ConfigKey:   configKey,
		ConfigValue: configValue,
	}
}
