package playerBLL

import (
	"fmt"
	"github.com/Jordanzuo/GameServerUtil_Go/manageUtil"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/globalBLL/playerNameBLL"
	"github.com/Jordanzuo/GameServer_Go/src/model/text/gameProperty"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/global/playerName"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
	"github.com/Jordanzuo/goutil/securityUtil"
	_ "github.com/Jordanzuo/goutil/stringUtil"
)

func init() {
	rpc.RegisterFunction(new(PlayerBLL))
}

type PlayerBLL int8

// 玩家登陆
// clientObj：客户端对象
// partnerId：合作商Id
// serverId：服务器Id
// gameVersionId：游戏版本Id
// resourceVersionId：资源版本Id
// name：玩家名称
// sign：登陆签名
func (playerBLL *PlayerBLL) C_Login(clientObj *rpc.Client, partnerId, serverId, gameVersionId, resourceVersionId int,
	name, sign string) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	// 判断合作商、服务器是否存在
	if manageUtil.IfServerExists(partnerId, serverId) == false {
		return rpc.GetResultStatusResponseObj(responseObj, rpc.ServerNotExists)
	}

	// 检测是否有游戏版本更新
	if gameVersionUrl, ok := manageUtil.IfHasNewGameVersion(partnerId, serverId, gameVersionId); ok {
		responseObj.Data = gameVersionUrl
		return rpc.GetResultStatusResponseObj(responseObj, rpc.NewGameVersion)
	}

	// 检测是否有资源版本更新
	if resourceVersionList, ok := manageUtil.IfHasNewResource(partnerId, gameVersionId, resourceVersionId); ok {
		responseObj.Data = resourceVersionList
		return rpc.GetResultStatusResponseObj(responseObj, rpc.NewResourceVersion)
	}

	// 验证签名是否正确
	if sign != securityUtil.Md5String(fmt.Sprintf("%s-%s", name, manageUtil.GetLoginKey(partnerId)), false) {
		return rpc.GetResultStatusResponseObj(responseObj, rpc.SignError)
	}

	// 判断玩家是否在缓存中已经存在
	var playerObj *player.Player
	var ok bool
	if playerObj, ok = GetPlayerByName(clientObj, name, partnerId, serverId); ok {
		// 判断是否重复登陆
		if oldClientObj, ok := rpc.GetClientByPlayerId(playerObj.Id); ok {
			// 如果不是同一个客户端，则先发送重复登陆的信息，然后玩家登出，客户端退出
			if clientObj != oldClientObj {
				// 先发送被路易下去的信息
				PushDataToClient(oldClientObj, rpc.GetResultStatusResponseObj(responseObj, rpc.LoginAgain))

				oldClientObj.LogoutAndQuit()
			}
		}

		// 更新玩家对象的ClientId
		playerObj.ClientId = clientObj.Id()
	} else {
		// todo 创建新玩家
		// playerObj = player.New(stringUtil.GetNewGUID(), name, partnerId, serverId, clientObj.Id())

		// // 注册玩家名称和Id
		// playerNameBLL.RegisterNameAndId(playerName.New(name, playerObj.Id()))
		return rpc.GetResultStatusResponseObj(responseObj, rpc.PlayerNotFound)
	}

	// 玩家上线
	clientObj.PlayerLogin(playerObj.Id, partnerId, serverId, gameVersionId, resourceVersionId)

	// 组装返回值
	responseObj.Data = assembleToClient(playerObj)

	return responseObj
}

// 修改玩家名称
// playerObj：玩家对象
// newName：新名称
func (playerBLL *PlayerBLL) C_AlterName(playerObj *player.Player, newName string) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	// 判断名字是否未改变
	if playerObj.Name == newName {
		return responseObj
	}

	// 判断新名称是否存在
	if _, exists := playerNameBLL.GetIdByName(newName); exists {
		return rpc.GetResultStatusResponseObj(responseObj, rpc.NameExists)
	}

	// 获取旧名称
	oldName := playerObj.Name

	// 修改名称
	playerObj.Name = newName

	// 移除旧名称，注册新名称
	playerNameBLL.UnRegisterNameAndId(playerName.New(oldName, playerObj.Id))
	playerNameBLL.RegisterNameAndId(playerName.New(playerObj.Name, playerObj.Id))

	// 推送信息begin

	// 组装data
	playerInfo := make(map[string]interface{})
	playerInfo[gameProperty.Name] = playerObj.Name

	data := make(map[string]interface{})
	data[gameProperty.PlayerInfo] = playerInfo

	pushResponseObj := rpc.GetInitResponseObj()
	pushResponseObj.Data = data

	PushDataToPlayer(playerObj, pushResponseObj)

	// 推送信息end

	return responseObj
}
