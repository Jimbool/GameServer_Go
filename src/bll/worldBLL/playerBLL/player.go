package playerBLL

import (
	_ "fmt"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/globalBLL/playerNameBLL"
	"github.com/Jordanzuo/GameServer_Go/src/dal/worldDAL/playerDAL"
	"github.com/Jordanzuo/GameServer_Go/src/model/text/gameProperty"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
	_ "github.com/Jordanzuo/goutil/logUtil"
)

// 根据玩家Id获得玩家对象
// clientObj：客户端对象
// id：玩家Id
// partnerId：合作商Id
// serverId：服务器Id
// 返回值：
// 玩家对象
// 是否存在玩家对象
func GetPlayerById(clientObj *rpc.Client, id string, partnerId, serverId int) (*player.Player, bool) {
	playerObj, exists := playerDAL.GetPlayerObj(id, partnerId, serverId)
	if exists {
		playerObj.ClientId = clientObj.Id()
	}

	return playerObj, exists
}

// 根据玩家名称获得玩家对象
// clientObj：客户端对象
// name：玩家名称
// partnerId：合作商Id
// serverId：服务器Id
// 返回值：
// 玩家对象
// 是否存在玩家对象
func GetPlayerByName(clientObj *rpc.Client, name string, partnerId, serverId int) (*player.Player, bool) {
	if id, ok := playerNameBLL.GetIdByName(name); ok {
		return GetPlayerById(clientObj, id, partnerId, serverId)
	}

	return nil, false
}

// 推送数据给客户端
// playerObj：玩家对象
// responseObj：结果对象
func PushDataToPlayer(playerObj *player.Player, responseObj rpc.ResponseObject) {
	if playerObj.ClientId > 0 {
		if clientObj, ok := rpc.GetClientById(playerObj.ClientId); ok {
			PushDataToClient(clientObj, responseObj)
		}
	}
}

// 推送数据给客户端
// playerObj：玩家对象
// responseObj：结果对象
func PushDataToClient(clientObj *rpc.Client, responseObj rpc.ResponseObject) {
	rpc.ResponseResult(clientObj, nil, responseObj)
}

func assembleToClient(playerObj *player.Player) map[string]interface{} {
	data := make(map[string]interface{})
	data[gameProperty.Id] = playerObj.Id
	data[gameProperty.Name] = playerObj.Name
	data[gameProperty.PartnerId] = playerObj.PartnerId
	data[gameProperty.ServerId] = playerObj.ServerId

	return data
}
