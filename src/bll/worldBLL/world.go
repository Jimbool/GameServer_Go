package worldBLL

// 使以下的包中的init方法自动调用，以实现自动初始化
import (
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/dynamicBLL"
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/globalBLL"
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/modelsBLL"
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/playerBLL"
)
