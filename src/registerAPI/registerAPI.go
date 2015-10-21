/*
用于让需要对外提供方法的包注册到RPC中而定义的包
*/
package registerAPI

// 为了让需要对外提供方法的对象将自身注册到RPC中而导入的包
import (
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/playerBLL"
	_ "github.com/Jordanzuo/GameServer_Go/src/test"
	// ...
)
