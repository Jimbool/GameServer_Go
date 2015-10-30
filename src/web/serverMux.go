package web

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/bll/worldBLL/modelsBLL/systemConfigBLL"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
	"github.com/Jordanzuo/goutil/logUtil"
	"net/http"
	"strings"
)

var (
	// 有效的ip映射表
	validIpMap = make(map[string]bool, 0)

	// 处理的方法映射表
	funcMap = make(map[string]func(http.ResponseWriter, *http.Request))
)

func init() {
	// 初始化有效的ip映射表
	allowIps := strings.Split(systemConfigBLL.AllowIps, ";")
	for _, ip := range allowIps {
		validIpMap[ip] = true
	}

	// 初始化处理的方法映射表
	funcMap["/API/charge"] = chargeCallback
	funcMap["/API/managecenter"] = managecenterCallback
}

// 定义自定义的Mux对象
type SelfDefineMux struct {
}

func (mux *SelfDefineMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 定义返回值
	responseObj := rpc.GetInitResponseObj()

	// 最终返回数据
	defer func() {
		// 只处理失败的情况；正确地情况已经转到各个具体的方法里面去处理了
		if responseObj.Code != rpc.Success {
			responseBytes, err := json.Marshal(responseObj)
			if err != nil {
				logUtil.Log(fmt.Sprintf("序列化输出结果%v出错", responseObj), logUtil.Error, true)
				return
			}

			fmt.Fprintln(w, string(responseBytes))
		}
	}()

	// 判断是否是POST方法
	if r.Method != "POST" {
		return
	}

	// 获取Ip
	ipAndPort := strings.Split(r.RemoteAddr, ":")
	if _, ok := validIpMap[ipAndPort[0]]; !ok {
		responseObj = rpc.GetResultStatusResponseObj(responseObj, rpc.IpNotAllowed)
		return
	}

	// 根据路径选择不同的处理方法
	if f, ok := funcMap[r.RequestURI]; ok {
		f(w, r)
	} else {
		http.NotFound(w, r)
	}
}
