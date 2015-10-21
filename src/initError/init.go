/*
初始化错误的包
*/
package initError

var (
	// 初始化的错误列表
	initErrorList []error = make([]error, 0, 32)
)

// 添加初始化错误
func RegisterInitError(err error) {
	initErrorList = append(initErrorList, err)
}

// 获取初始化错误列表
func GetInitErrorList() []error {
	return initErrorList
}
