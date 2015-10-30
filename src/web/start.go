package web

import (
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/config"
	"github.com/Jordanzuo/goutil/logUtil"
	"net/http"
	"sync"
)

// 启动服务器
// wg：WaitGroup
func StartServer(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	// 设置访问的路由
	mux := new(SelfDefineMux)

	// 启动Web服务器监听
	err := http.ListenAndServe(config.WebServerAddress, mux)
	if err != nil {
		msg := fmt.Sprintf("ListenAndServe失败，错误信息为：%s", err)
		fmt.Println(msg)
		logUtil.Log(msg, logUtil.Error, true)
	}
}
