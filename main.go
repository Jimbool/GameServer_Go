package main

// 逻辑本身需要用到的包
import (
	"fmt"
	"github.com/Jordanzuo/GameServerUtil_Go/manageUtil"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/dynamicBLL/serverGroupSettingBLL"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/modelsBLL/systemConfigBLL"
	"github.com/Jordanzuo/GameServer_Go/src/initError"
	rpc "github.com/Jordanzuo/RPCServer_Go"
	"github.com/Jordanzuo/goutil/logUtil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

// 只是用于初始化数据的import
import (
	// 为了让需要对外提供方法的对象将自身注册到RPC中而导入的包
	_ "github.com/Jordanzuo/GameServer_Go/src/registerAPI"

	// 用于让所有的模型数据进行初始化，并进行数据验证
	_ "github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL"
)

const (
	// 日志文件路径后缀
	LOG_PATH_SUFFIX = "LOG"
)

var (
	wg sync.WaitGroup
)

func init() {
	// 设置日志文件的存储目录
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	logUtil.SetLogPath(filepath.Join(logPath, LOG_PATH_SUFFIX))

	// 设置WaitGroup需要等待的数量
	wg.Add(1)
}

// 处理系统信号
func signalProc() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	for {
		// 准备接收信息
		<-sigs

		// 一旦收到信号，则表明管理员希望退出程序，则先保存信息，然后退出

		// todo：执行一些收尾的工作

		// 退出
		os.Exit(0)
	}
}

func main() {
	// 处理系统信号
	go signalProc()

	// 检查初始化错误
	errs := initError.GetInitErrorList()
	// if len(errs) > 0 {
	// 	fmt.Println("初始化游戏世界错误，错误信息为：")
	// 	for index, err := range errs {
	// 		fmt.Println(index+1, ":", err)
	// 	}
	// 	os.Exit(0)
	// }
	_ = errs

	// 激活服务器
	err := manageUtil.ActivateServer(serverGroupSettingBLL.GetServerGroupSettingConfig().ManageCenterUrl(), serverGroupSettingBLL.GetServerGroupSettingConfig().ServerGroupId())
	if err != nil {
		fmt.Println("激活服务器失败，错误信息为：", err)
		os.Exit(0)
	}

	// 设置RPC服务器所需参数
	rpc.SetRPCParam(manageUtil.GetServerUrl(), systemConfigBLL.CheckExpiredInterval, systemConfigBLL.ClientExpiredSeconds)

	// 启动Socket服务器
	go rpc.StartServer(&wg)

	// todo：启动Web服务器

	// 阻塞等待，以免main线程退出
	wg.Wait()
}
