package rpc

// 服务端响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

// 返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) String() string {
	if message, ok := status[rs]; ok {
		return message
	}

	return ""
}

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Success ResultStatus = 0

	// 无效的用户自定义编号
	InvalidUserDefineCode ResultStatus = -1

	// 数据格式错误
	DataFormatError ResultStatus = -2

	// 未找到目标方法
	NoTargetMethod ResultStatus = -3

	// 参数不匹配
	ParamNotMatch ResultStatus = -4

	// 数据错误
	DataError ResultStatus = -5

	// 玩家尚未登陆
	PlayerNotLogin ResultStatus = -6

	// 未找到玩家对象
	PlayerNotFound ResultStatus = -7

	// 重复登陆
	LoginAgain ResultStatus = -1101

	// 签名错误
	SignError ResultStatus = -1102

	// 名称已经存在
	NameExists ResultStatus = -1103

	// 服务器不存在
	ServerNotExists ResultStatus = -1104

	// 有新版本
	NewGameVersion ResultStatus = -1105

	// 有新的资源版本
	NewResourceVersion ResultStatus = -1106
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status map[ResultStatus]string = map[ResultStatus]string{
	0:     "Success",
	-1:    "InvalidUserDefineCode",
	-2:    "DataFormatError",
	-3:    "NoTargetMethod",
	-4:    "ParamNotMatch",
	-5:    "DataError",
	-6:    "PlayerNotLogin",
	-7:    "PlayerNotFound",
	-1101: "LoginAgain",
	-1102: "SignError",
	-1103: "NameExists",
	-1104: "ServerNotExists",
	-1105: "NewGameVersion",
	-1106: "NewResourceVersion",
}
