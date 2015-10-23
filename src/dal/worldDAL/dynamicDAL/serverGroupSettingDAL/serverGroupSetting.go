/*
服务器组设置数据处理层
*/
package serverGroupSettingDAL

import (
	"database/sql"
	"github.com/Jordanzuo/GameServer_Go/src/dal"
)

var (
	getListCommand string = "SELECT ServerGroupId, ManageCenterUrl FROM d_server_group_setting_c WHERE ServerGroupId = ?;"
)

func GetList(serverGroupId int) (*sql.Rows, error) {
	db := dal.GameModelDB()
	rows, err := db.Query(getListCommand, serverGroupId)

	return rows, err
}
