/*
玩家名称的数据处理包
*/
package playerNameDAL

import (
	"database/sql"
	"github.com/Jordanzuo/GameServer_Go/src/dal"
)

var (
	getListCommand string = "SELECT Name, Id FROM p_player;"
)

func GetList() (*sql.Rows, error) {
	// for test
	return nil, nil

	db := dal.GameDB()
	rows, err := db.Query(getListCommand)

	return rows, err
}
