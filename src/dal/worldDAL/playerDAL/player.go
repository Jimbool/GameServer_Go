package playerDAL

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/dal"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/garyburd/redigo/redis"
)

var (
	getListCommand string = "SELECT Id, Name, PartnerId, ServerId FROM p_player WHERE Id = ? AND PartnerId = ? AND ServerId = ?;"
)

func getPlayerFromRedis(id string, partnerId, serverId int) (*player.Player, bool) {
	var playerObj *player.Player

	// 获取redis连接对象
	conn := dal.GetRedisConn()
	defer conn.Close()

	// 执行GET命令
	returnValue, err := conn.Do("GET", id)
	dal.HandleRedisError(err)

	// 解析返回值
	if returnValue != nil {
		playerBytes, err := redis.Bytes(returnValue, err)
		dal.HandleRedisError(err)

		if playerBytes != nil {
			json.Unmarshal(playerBytes, &playerObj)

			return playerObj, true
		}
	}

	return nil, false
}

func getPlayerFromMysql(id string, partnerId, serverId int) (*player.Player, bool) {
	var playerObj *player.Player

	// 获取mysql数据库连接
	db := dal.GameDB()

	// 执行查询命令
	rows, err := db.Query(getListCommand, id, partnerId, serverId)
	dal.HandleMysqlError(err)

	// 解析数据
	for rows.Next() {
		var name string
		err := rows.Scan(&id, &name, &partnerId, &serverId)
		dal.HandleMysqlError(err)

		playerObj = player.New(id, name, partnerId, serverId, 0)
	}

	return playerObj, playerObj != nil
}

// 获取玩家对象
// id：玩家id
// partnerId：合作商Id
// serverId：服务器Id
// 返回值
// 玩家对象
// 是否找到玩家对象
func GetPlayerObj(id string, partnerId, serverId int) (*player.Player, bool) {
	var playerObj *player.Player
	var exists bool

	// 先从redis中查找
	playerObj, exists = getPlayerFromRedis(id, partnerId, serverId)
	if exists {
		return playerObj, exists
	}

	fmt.Println("在redis中没有找到，从mysql中进行查找")

	// 否则从mysql中查找
	playerObj, exists = getPlayerFromMysql(id, partnerId, serverId)

	// 如果取到数据则保存到redis中
	if exists {
		SavePlayerToRedis(playerObj)
	}

	return playerObj, exists
}

// 保存玩家到Redis中
func SavePlayerToRedis(playerObj *player.Player) {
	// 获取redis连接对象
	conn := dal.GetRedisConn()
	defer conn.Close()

	// 讲玩家对象序列化
	jsonBytes, err := json.Marshal(playerObj)
	if err != nil {
		dal.HandleOtherError(err)
	}

	// 执行保存数据的redis命令
	_, err = conn.Do("SET", playerObj.Id, jsonBytes)
	dal.HandleRedisError(err)

	// 设置过期时间:30m
	_, err = conn.Do("EXPIRE", playerObj.Id, 30*60)
	dal.HandleRedisError(err)
}
