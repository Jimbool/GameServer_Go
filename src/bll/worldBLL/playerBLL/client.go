package playerBLL

import (
	"fmt"
	"github.com/Jordanzuo/GameServerUtil_Go/manageUtil"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/globalBLL/playerNameBLL"
	"github.com/Jordanzuo/GameServer_Go/src/model/text/gameProperty"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/global/playerName"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/Jordanzuo/GameServer_Go/src/resultStatus"
	rpc "github.com/Jordanzuo/RPCServer_Go"
	"github.com/Jordanzuo/goutil/securityUtil"
	"github.com/Jordanzuo/goutil/stringUtil"
)

func init() {
	rpc.RegisterFunction(new(PlayerBLL))
}

type PlayerBLL int8

// 玩家登陆
// clientObj：客户端对象
// partnerId：合作商Id
// serverId：服务器Id
// name：玩家名称
// sign：登陆签名
func (playerBLL *PlayerBLL) C_Login(clientObj *rpc.Client, partnerId, serverId int, name, sign string) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	// 验证签名是否正确
	if sign != securityUtil.Md5String(fmt.Sprintf("%s-%s", name, manageUtil.GetLoginKey(partnerId)), false) {
		return resultStatus.GetResponseObject(responseObj, resultStatus.SignError)
	}

	// 判断玩家是否在缓存中已经存在
	var playerObj rpc.IPlayer
	var ok bool
	if playerObj, ok = GetPlayerByName(name); ok {
		// 判断是否重复登陆
		if oldClientObj, ok := rpc.GetClientById(playerObj.ClientId()); ok {
			// 先发送被路易下去的信息
			rpc.ResponseResult(oldClientObj, nil, resultStatus.GetResponseObject(responseObj, resultStatus.LoginAgain))

			// 如果不是同一个客户端，则玩家登出，客户端退出
			if clientObj != oldClientObj {
				oldClientObj.LogoutAndQuit()
			}
		}

		// 更新玩家对象的ClientId
		playerObj.SetClientId(clientObj.Id())
	} else {
		playerObj = player.New(stringUtil.GetNewGUID(), name, clientObj.Id(), partnerId, serverId)

		// 注册玩家名称和Id
		playerNameBLL.RegisterNameAndId(playerName.New(name, playerObj.Id()))
	}

	// 将玩家添加到列表中
	rpc.RegisterPlayer(rpc.NewClientPlayerPair(clientObj, playerObj))

	// 将playerObj转化为Player对象
	newPlayerObj := GetPlayerFromIPlayer(playerObj)

	// 组装返回值
	data := make(map[string]interface{})
	data[gameProperty.Id] = newPlayerObj.Id()
	data[gameProperty.Name] = newPlayerObj.Name()
	data[gameProperty.PartnerId] = newPlayerObj.PartnerId()
	data[gameProperty.ServerId] = newPlayerObj.ServerId()

	responseObj.Data = data

	return responseObj
}

// 修改玩家名称
// iplayer：玩家对象
// newName：新名称
func (playerBLL *PlayerBLL) C_AlterName(iplayer rpc.IPlayer, newName string) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	// 获取具体的玩家对象
	playerObj := GetPlayerFromIPlayer(iplayer)

	// 判断名字是否未改变
	if playerObj.Name() == newName {
		return responseObj
	}

	// 判断新名称是否存在
	if _, exists := playerNameBLL.GetIdByName(newName); exists {
		return resultStatus.GetResponseObject(responseObj, resultStatus.NameExists)
	}

	// 获取旧名称
	oldName := playerObj.Name()

	// 修改名称
	playerObj.SetName(newName)

	// 移除旧名称，注册新名称
	playerNameBLL.UnRegisterNameAndId(playerName.New(oldName, playerObj.Id()))
	playerNameBLL.RegisterNameAndId(playerName.New(playerObj.Name(), playerObj.Id()))

	// 推送信息
	if clientObj, ok := rpc.GetClientByPlayer(playerObj); ok {
		targetResponseObj := rpc.GetInitResponseObj()

		// 组装data
		playerInfo := make(map[string]interface{})
		playerInfo[gameProperty.Name] = playerObj.Name()

		data := make(map[string]interface{})
		data[gameProperty.PlayerInfo] = playerInfo

		targetResponseObj.Data = data

		rpc.ResponseResult(clientObj, nil, targetResponseObj)
	}

	return responseObj
}
