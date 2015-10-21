package playerBLL

import (
	"errors"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/globalBLL/playerNameBLL"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	rpc "github.com/Jordanzuo/RPCServer_Go"
)

// 获取玩家列表
func GetPlayerList() map[string]rpc.IPlayer {
	return rpc.PlayerList()
}

// 根据玩家Id获得玩家对象
// id：玩家Id
// 返回值：
// 玩家对象
// 是否存在玩家对象
func GetPlayerById(id string) (rpc.IPlayer, bool) {
	if playerObj, ok := GetPlayerList()[id]; ok {
		return playerObj, true
	}

	return nil, false
}

// 根据玩家名称获得玩家对象
// id：玩家Id
// 返回值：
// 玩家对象
// 是否存在玩家对象
func GetPlayerByName(name string) (rpc.IPlayer, bool) {
	if id, ok := playerNameBLL.GetIdByName(name); ok {
		return GetPlayerById(id)
	}

	return nil, false
}

// 从IPlayer中得到玩家对象
// playerObj：IPlayer的玩家对象
// 返回值：
// 玩家对象
func GetPlayerFromIPlayer(playerObj rpc.IPlayer) *player.Player {
	if newPlayer, ok := playerObj.(*player.Player); ok {
		return newPlayer
	}

	panic(errors.New("IPlayer对象转换为player.Player失败"))
}
