package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
	"time"
)

var (
	// 客户端连接列表
	clientList = make(map[int32]*Client)

	// 定义增加、删除客户端channel；
	clientAddChan    = make(chan *Client)
	clientRemoveChan = make(chan *Client, 50)

	// 玩家登陆的模块名称
	mPlayerLoginModuleName = "Player"

	// 玩家登陆的方法名称
	mPlayerLoginMethodName = "Login"
)

// 设置玩家登陆参数
// playerLoginModuleName：玩家登陆的模块名称
// playerLoginMethodName：玩家登陆的方法名称
func SetPlayerLoginParam(playerLoginModuleName, playerLoginMethodName string) {
	mPlayerLoginModuleName = playerLoginModuleName
	mPlayerLoginMethodName = playerLoginMethodName
}

func init() {
	// 启动处理增加、删除客户端channel；增加、删除玩家的channel的gorountine
	go handleChannel()
}

// 获取客户端列表
func ClientList() map[int32]*Client {
	return clientList
}

func clearExpiredClient() {
	// 处理内部未处理的异常，以免导致主线程退出，从而导致系统崩溃
	defer func() {
		if r := recover(); r != nil {
			logUtil.LogUnknownError(r)
		}
	}()

	for {
		// 休眠指定的时间（单位：秒）(放在此处是因为程序刚启动时并没有过期的客户端，所以先不用占用资源；)
		time.Sleep(CheckExpiredInterval() * time.Second)

		// 清理之前的客户端数量和玩家数量
		beforeClientCount := len(clientList)

		// 获取本次清理的客户端数量
		expiredClientCount := 0

		// 开始清理
		for _, item := range clientList {
			if item.HasExpired() {
				expiredClientCount++
				item.Quit()
			}
		}

		// 记录日志
		if expiredClientCount > 0 {
			logUtil.Log(fmt.Sprintf("清理前的客户端数量为：%d，本次清理不活跃的数量为：%d", beforeClientCount, expiredClientCount), logUtil.Debug, true)
		}
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
		case clientObj := <-clientAddChan:
			addClient(clientObj)
		case clientObj := <-clientRemoveChan:
			removeClient(clientObj)
		default:
			// 休眠一下，防止CPU过高
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 添加新的客户端
// clientObj：客户端对象
func registerClient(clientObj *Client) {
	clientAddChan <- clientObj
}

// 移除客户端
// clientObj：客户端对象
func unRegisterClient(clientObj *Client) {
	clientRemoveChan <- clientObj
}

// 添加一个新的客户端对象到列表中
// clientObj：客户端对象
func addClient(clientObj *Client) {
	clientList[clientObj.Id()] = clientObj
}

// 移除一个客户端对象
// clientObj：客户端对象
func removeClient(clientObj *Client) {
	delete(clientList, clientObj.Id())
}

// 根据玩家对象获取对应的客户端对象
// id：客户端Id
// 返回值：
// 客户端对象
// 是否存在客户端对象
func GetClientById(id int32) (*Client, bool) {
	if clientObj, ok := clientList[id]; ok {
		return clientObj, true
	}

	return nil, false
}

// 根据玩家Id获得客户端对象
// playerId：玩家Id
// 返回值：
// 客户端对象
// 是否存在客户端对象
func GetClientByPlayerId(playerId string) (*Client, bool) {
	for _, clientObj := range clientList {
		if clientObj.PlayerId() == playerId {
			return clientObj, true
		}
	}

	return nil, false
}

// 处理请求
// clientObj：对应的客户端对象
// id：客户端请求唯一标识
// request：请求内容字节数组(json格式)
// 返回值：无
func handleRequest(clientObj *Client, id int, request []byte) {
	responseObj := GetInitResponseObj()

	// 解析请求字符串
	var requestObj *RequestObject

	// 提取请求内容
	err := json.Unmarshal(request, &requestObj)
	if err != nil {
		logUtil.Log(fmt.Sprintf("反序列化%s出错，错误信息为：%s", string(request), err), logUtil.Error, true)
		ResponseResult(clientObj, requestObj, GetResultStatusResponseObj(responseObj, DataFormatError))
		return
	}

	// 对requestObj的属性Id赋值
	requestObj.Id = id

	// 对参数要特殊处理，将playerObj或clientObj添加到最前面
	parameters := make([]interface{}, 0)
	if requestObj.ModuleName == mPlayerLoginModuleName && requestObj.MethodName == mPlayerLoginMethodName {
		parameters = append(parameters, interface{}(clientObj))
		parameters = append(parameters, requestObj.Parameters...)
	} else {
		// 判断玩家是否已经登陆
		if clientObj.PlayerId() == "" {
			ResponseResult(clientObj, requestObj, GetResultStatusResponseObj(responseObj, PlayerNotLogin))
			return
		}

		// 判断是否能找到玩家
		f := GetPlayerFunc()
		if playerObj, ok := f(clientObj, clientObj.PlayerId(), clientObj.PartnerId(), clientObj.ServerId()); !ok {
			ResponseResult(clientObj, requestObj, GetResultStatusResponseObj(responseObj, PlayerNotFound))
			return
		} else {
			parameters = append(parameters, interface{}(playerObj))
			parameters = append(parameters, requestObj.Parameters...)
		}
	}
	requestObj.Parameters = parameters

	// 调用方法
	callFunction(clientObj, requestObj)
}
