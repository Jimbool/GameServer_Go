package web

import (
	"fmt"
	"github.com/Jordanzuo/GameServerUtil_Go/manageUtil"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/playerBLL"
	"github.com/Jordanzuo/GameServer_Go/src/config"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
	"io/ioutil"
	"net/http"
)

func managecenterCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to visit managecenter callback method")
	} else {
		// 接受数据
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Fail")
			return
		}

		fmt.Println("Content:", string(content))

		// 解析数据
		err = manageUtil.ParseData(config.ServerGroupId, string(content))
		if err != nil {
			fmt.Fprintf(w, "Fail")
			return
		}

		// 检查服务器是否维护，是否有新的游戏版本、资源版本
		go checkData()

		fmt.Fprintf(w, "Success")
	}
}

func checkData() {
	responseObj := rpc.GetInitResponseObj()

	// 判断服务器是否维护
	if manageUtil.IfServerMaintain() {
		responseObj = rpc.GetResultStatusResponseObj(responseObj, rpc.ServerMaintain)

		// 向所有玩家推送信息，并断开连接
		for _, clientObj := range rpc.ClientList() {
			playerBLL.PushDataToClient(clientObj, responseObj)
			clientObj.LogoutAndQuit()
		}

		// 如果服务器维护，则所有人都退出；无需再判断游戏版本、资源版本的更新
		return
	}

	// 判断是否有新版本
	for _, clientObj := range rpc.ClientList() {
		// 判断是否有新版本
		if gameVersionUrl, ok := manageUtil.IfHasNewGameVersion(clientObj.PartnerId(), clientObj.ServerId(), clientObj.GameVersionId()); ok {
			responseObj = rpc.GetResultStatusResponseObj(responseObj, rpc.NewGameVersion)
			responseObj.Data = gameVersionUrl

			// 推送信息并断开连接
			playerBLL.PushDataToClient(clientObj, responseObj)
			clientObj.LogoutAndQuit()

			// 如果有游戏版本更新则不判断资源更新
			continue
		}

		// 判断是否有新资源
		if resourceVersionList, ok := manageUtil.IfHasNewResource(clientObj.PartnerId(), clientObj.GameVersionId(), clientObj.ResourceVersionId()); ok {
			responseObj = rpc.GetResultStatusResponseObj(responseObj, rpc.NewResourceVersion)
			responseObj.Data = resourceVersionList

			// 推送信息并断开连接
			playerBLL.PushDataToClient(clientObj, responseObj)
			clientObj.LogoutAndQuit()
		}
	}
}
