/*
项目配置的逻辑处理包，初始化所有的配置内容，其它代码需要配置时都从此包内来获取
*/
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const (
	// 配置文件名称
	CONFIG_FILE_NAME = "config.ini"
)

var (
	// 游戏模型数据库连接字符串
	GameModelDBConnection string

	// 游戏数据库连接字符串
	GameDBConnection string

	// 服务器组Id
	ServerGroupId int
)

func init() {
	// 由于服务器的运行依赖于init中执行的逻辑，所以如果出现任何的错误都直接panic，让程序启动失败；而不是让它启动成功，但是在运行时出现错误

	// 读取配置文件（一次性读取整个文件，则使用ioutil）
	bytes, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		panic(err)
	}

	// 使用json反序列化
	config := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}

	// 解析GameModelDBConnection
	gameModelDBConnection, ok := config["GameModelDBConnection"]
	if !ok {
		panic(errors.New("不存在名为GameModelDBConnection的配置或配置为空"))
	}
	gameModelDBConnection_string, ok := gameModelDBConnection.(string)
	if !ok {
		panic(errors.New("GameModelDBConnection必须为字符串类型"))
	}

	// 设置GameModelDBConnection
	GameModelDBConnection = gameModelDBConnection_string

	// 解析GameDBConnection
	gameDBConnection, ok := config["GameDBConnection"]
	if !ok {
		panic(errors.New("不存在名为GameDBConnection的配置或配置为空"))
	}
	gameDBConnection_string, ok := gameDBConnection.(string)
	if !ok {
		panic(errors.New("GameDBConnection必须为字符串类型"))
	}

	// 设置GameDBConnection
	GameDBConnection = gameDBConnection_string

	// 解析ServerGroupId
	serverGroupId, ok := config["ServerGroupId"]
	if !ok {
		panic(errors.New("不存在名为ServerGroupId的配置或配置为空"))
	}

	serverGroupId_float64, ok := serverGroupId.(float64)
	if !ok {
		panic(errors.New("ServerGroupId必须是int类型"))
	}

	// 设置ServerGroupId
	ServerGroupId = int(serverGroupId_float64)
}
