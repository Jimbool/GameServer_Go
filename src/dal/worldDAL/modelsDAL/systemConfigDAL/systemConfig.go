package systemConfigDAL

import (
	"database/sql"
	"github.com/Jordanzuo/GameServer_Go/src/dal"
)

var (
	getListCommand string = "SELECT ConfigKey, ConfigValue FROM b_system_c;"
)

func GetList() (*sql.Rows, error) {
	db := dal.GameModelDB()
	rows, err := db.Query(getListCommand)

	return rows, err
}
