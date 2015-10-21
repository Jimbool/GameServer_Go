/*
定义方法的返回值状态的包
*/
package resultStatus

import (
	"errors"
	rpc "github.com/Jordanzuo/RPCServer_Go"
)

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
// 范围-1101~-9999
const (
	LoginAgain = -1101
	SignError  = -1102
	NameExists = -1103
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status map[int]string = map[int]string{
	-1101: "LoginAgain",
	-1102: "SignError",
	-1103: "NameExists",
}

// 获取用户自定义的返回对象
func GetResponseObject(responseObj rpc.ResponseObject, code int) rpc.ResponseObject {
	var message string
	var ok bool
	if message, ok = status[code]; !ok {
		panic(errors.New("未定义的code，请检查resultStatus.go文件"))
	}

	return rpc.GetUserDefineResponseObj(responseObj, code, message)
}
