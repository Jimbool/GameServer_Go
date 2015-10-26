/*
服务器组设置逻辑包
*/
package serverGroupSettingBLL

import (
	"github.com/Jordanzuo/GameServer_Go/src/config"
	"github.com/Jordanzuo/GameServer_Go/src/dal/worldDAL/dynamicDAL/serverGroupSettingDAL"
	"github.com/Jordanzuo/GameServer_Go/src/initError"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/dynamic/serverGroupSetting"
)

var (
	serverGroupSettingConfig *serverGroupSetting.ServerGroupSettingConfig
)

func init() {
	rows, err := serverGroupSettingDAL.GetList(config.ServerGroupId)
	if err != nil {
		initError.RegisterInitError(err)
		return
	}

	for rows.Next() {
		var serverGroupId int
		var manageCenterUrl string
		err := rows.Scan(&serverGroupId, &manageCenterUrl)
		if err != nil {
			initError.RegisterInitError(err)
			continue
		}

		// 构造对象
		serverGroupSettingConfig = serverGroupSetting.New(serverGroupId, manageCenterUrl)
	}
}

// 获取服务器组设置配置对象
// 返回值：
// 服务器组设置配置对象
func GetServerGroupSettingConfig() *serverGroupSetting.ServerGroupSettingConfig {
	return serverGroupSettingConfig
}
