/*
处理玩家名称的包
*/
package playerNameBLL

import (
	"github.com/Jordanzuo/GameServer_Go/src/dal/worldDAL/globalDAL/playerNameDAL"
	"github.com/Jordanzuo/GameServer_Go/src/initError"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/global/playerName"
	"github.com/Jordanzuo/goutil/logUtil"
	"time"
)

var (
	playerNameList = make(map[string]*playerName.NameAndId)

	// 定义增加、移除玩家名称的通道
	playerNameAddChan    = make(chan *playerName.NameAndId)
	playerNameRemoveChan = make(chan *playerName.NameAndId, 50)
)

func init() {
	// 启动处理增加、删除玩家名称channel的gorountine
	go handleChannel()
}

func init() {
	// 加载数据
	rows, err := playerNameDAL.GetList()
	if err != nil {
		initError.RegisterInitError(err)
		return
	}

	for rows.Next() {
		var name string
		var id string
		err := rows.Scan(&name, &id)
		if err != nil {
			initError.RegisterInitError(err)
			continue
		}

		// 添加到列表中
		addNameAndId(playerName.New(name, id))
	}
}

func handleChannel() {
	// 处理内部未处理的异常，以免导致主线程退出，从而导致系统崩溃
	defer func() {
		if r := recover(); r != nil {
			logUtil.LogUnknownError(r)
		}
	}()

	for {
		select {
		case nameAndIdObj := <-playerNameAddChan:
			addNameAndId(nameAndIdObj)
		case nameAndIdObj := <-playerNameRemoveChan:
			removeNameAndId(nameAndIdObj)
		default:
			// 休眠一下，防止CPU过高
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func addNameAndId(nameAndIdObj *playerName.NameAndId) {
	playerNameList[nameAndIdObj.Name] = nameAndIdObj
}

func removeNameAndId(nameAndIdObj *playerName.NameAndId) {
	delete(playerNameList, nameAndIdObj.Name)
}

// 注册玩家名称和Id的映射
// nameAndId：名称和Id对应对象
func RegisterNameAndId(nameAndId *playerName.NameAndId) {
	playerNameAddChan <- nameAndId
}

// 取消玩家名称和Id的映射的注册
// nameAndId：名称和Id对应对象
func UnRegisterNameAndId(nameAndId *playerName.NameAndId) {
	playerNameRemoveChan <- nameAndId
}

// 根据玩家名称获得Id
func GetIdByName(name string) (string, bool) {
	if nameAndIdObj, ok := playerNameList[name]; ok {
		return nameAndIdObj.Id, true
	}

	return "", false
}
