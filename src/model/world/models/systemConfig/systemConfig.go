package systemConfig

// 系统配置项
type SystemConfigItem struct {
	// 配置的Key
	configKey string

	// 配置的值
	configValue string
}

func New(configKey, configValue string) *SystemConfigItem {
	return &SystemConfigItem{
		configKey:   configKey,
		configValue: configValue,
	}
}

func (item *SystemConfigItem) ConfigKey() string {
	return item.configKey
}

func (item *SystemConfigItem) ConfigValue() string {
	return item.configValue
}
