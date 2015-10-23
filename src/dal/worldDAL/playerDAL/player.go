package playerDAL

import (
	"database/sql"
	"github.com/Jordanzuo/GameServer_Go/src/dal"
)

var (
	getListCommand string = "SELECT Id, Name, PartnerId, ServerId FROM p_player WHERE Id = ? AND PartnerId = ? AND ServerId = ?;"
)

func GetList(id string, partnerId, serverId int) (*sql.Rows, error) {
	db := dal.GameDB()
	rows, err := db.Query(getListCommand, id, partnerId, serverId)

	return rows, err
}
