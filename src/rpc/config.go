package rpc

import (
	"errors"
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"time"
)

var (
	// 服务器监听地址
	mServerAddress string

	// 检测客户端过期的时间间隔（单位：秒）
	mCheckExpiredInterval time.Duration = time.Duration(0)

	// 客户端过期的秒数
	mClientExpiredSeconds time.Duration = time.Duration(0)

	// 获取玩家对象的方法
	mGetPlayerFunc func(*Client, string, int, int) (*player.Player, bool)
)

// 设置RPC服务器所需参数
// serverAddress：服务器监听地址
// checkExpiredInterval：检测客户端过期的时间间隔（单位：秒），建议大于或等于60秒
// clientExpiredSeconds：客户端过期的秒数，建议大于或等于60秒
// getPlayerFunc：获取玩家对象的方法
func SetRPCParam(serverAddress string, checkExpiredInterval, clientExpiredSeconds time.Duration, getPlayerFunc func(*Client, string, int, int) (*player.Player, bool)) {
	mServerAddress = serverAddress
	mCheckExpiredInterval = checkExpiredInterval
	mClientExpiredSeconds = clientExpiredSeconds
	mGetPlayerFunc = getPlayerFunc
}

// 获取服务器监听地址
// 返回值：
// 服务器监听地址
func ServerAddress() string {
	if mServerAddress == "" {
		panic(errors.New("mServerAddress尚未设置，请先设置"))
	}

	return mServerAddress
}

// 获取检测客户端过期的时间间隔（单位：秒）
// 返回值：
// 检测客户端过期的时间间隔（单位：秒）
func CheckExpiredInterval() time.Duration {
	if mCheckExpiredInterval == time.Duration(0) {
		panic(errors.New("mCheckExpiredInterval尚未设置，请先设置"))
	}

	return mCheckExpiredInterval
}

// 获取客户端过期的秒数
// 返回值：
// 客户端过期的秒数
func ClientExpiredSeconds() time.Duration {
	if mClientExpiredSeconds == time.Duration(0) {
		panic(errors.New("mClientExpiredSeconds尚未设置，请先设置"))
	}

	return mClientExpiredSeconds
}

// 获取玩家对象的方法
func GetPlayerFunc() func(*Client, string, int, int) (*player.Player, bool) {
	if mGetPlayerFunc == nil {
		panic(errors.New("获取玩家对象的方法尚未设置，请先进行设置"))
	}

	return mGetPlayerFunc
}
