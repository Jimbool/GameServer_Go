package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
)

// 获取初始的响应对象
// 返回值：
// 响应对象
func GetInitResponseObj() ResponseObject {
	return ResponseObject{
		Code:    Success,
		Message: "",
		Data:    nil,
	}
}

// 获取响应类型为数据错误的响应对象
// 返回值：
// 响应对象
func GetDataErrorReponseObj() ResponseObject {
	return ResponseObject{
		Code:    DataError,
		Message: DataError.String(),
		Data:    nil,
	}
}

// 获取指定响应类型的响应对象
// responseObj：响应对象
// rs：响应类型对象
// 返回值：
// 响应对象
func GetResultStatusResponseObj(responseObj ResponseObject, rs ResultStatus) ResponseObject {
	responseObj.Code = rs
	responseObj.Message = rs.String()

	return responseObj
}

// 发送响应结果
// clientObj：客户端对象
// requestObj：请求对象（如果为nil则代表是服务端主动推送信息，否则为客户端请求信息）
// responseObject：响应对象（不能为指针类型，否则在RegisterFunction时判断类型会出错）
func ResponseResult(clientObj *Client, requestObj *RequestObject, responseObj interface{}) {
	b, err := json.Marshal(responseObj)
	if err != nil {
		logUtil.Log(fmt.Sprintf("序列化输出结果%v出错", responseObj), logUtil.Error, true)
		return
	}

	if requestObj == nil {
		clientObj.SendByteMessage(0, b)
	} else {
		clientObj.SendByteMessage(requestObj.Id, b)
	}
}
