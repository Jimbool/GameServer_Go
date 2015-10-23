package rpc

import (
	"encoding/binary"
	"github.com/Jordanzuo/goutil/intAndBytesUtil"
	"net"
	"sync/atomic"
	"time"
)

const (
	// 包头的长度
	HEADER_LENGTH = 4

	// 定义请求、响应数据的前缀的长度
	ID_LENGTH = 4
)

var (
	// 字节的大小端顺序
	byterOrder = binary.LittleEndian

	// 全局客户端的id，从1开始进行自增
	globalClientId int32 = 0
)

// 获得自增的id值
func getIncrementId() int32 {
	atomic.AddInt32(&globalClientId, 1)

	return globalClientId
}

// 定义客户端对象，以实现对客户端连接的封装
type Client struct {
	// 唯一标识
	id int32

	// 客户端连接对象
	conn net.Conn

	// 接收到的消息内容
	content []byte

	// 上次活跃时间
	activeTime time.Time

	// 玩家id
	playerId string

	// 合作商Id
	partnerId int

	// 服务器Id
	serverId int

	// 游戏版本号
	gameVersionId int

	// 资源版本号
	resourceVersionId int
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func NewClient(conn net.Conn) *Client {
	return &Client{
		id:         getIncrementId(),
		conn:       conn,
		content:    make([]byte, 0, 1024),
		activeTime: time.Now(),
		// 与玩家相关的属性使用默认值
	}
}

// 获取唯一标识
func (c *Client) Id() int32 {
	return c.id
}

// 获取玩家Id
func (c *Client) PlayerId() string {
	return c.playerId
}

// 获取合作商Id
func (c *Client) PartnerId() int {
	return c.partnerId
}

// 获取服务器Id
func (c *Client) ServerId() int {
	return c.serverId
}

// 获取游戏版本Id
func (c *Client) GameVersionId() int {
	return c.gameVersionId
}

// 获取资源版本Id
func (c *Client) ResourceVersionId() int {
	return c.resourceVersionId
}

// 追加内容
// content：新的内容
// 返回值：无
func (c *Client) AppendContent(content []byte) {
	c.content = append(c.content, content...)
	c.activeTime = time.Now()
}

// 获取有效的消息
// 返回值：
// 消息对应客户端的唯一标识
// 消息内容
// 是否含有有效数据
func (c *Client) GetValieMessage() (int, []byte, bool) {
	// 判断是否包含头部信息
	if len(c.content) < HEADER_LENGTH {
		return 0, nil, false
	}

	// 获取头部信息
	header := c.content[:HEADER_LENGTH]

	// 将头部数据转换为内部的长度
	contentLength := intAndBytesUtil.BytesToInt(header, byterOrder)

	// 判断长度是否满足
	if len(c.content)-HEADER_LENGTH < contentLength {
		return 0, nil, false
	}

	// 提取消息内容
	content := c.content[HEADER_LENGTH : HEADER_LENGTH+contentLength]

	// 将对应的数据截断，以得到新的数据
	c.content = c.content[HEADER_LENGTH+contentLength:]

	// 截取内容的前4位
	idBytes, content := content[:ID_LENGTH], content[ID_LENGTH:]

	// 提取id
	id := intAndBytesUtil.BytesToInt(idBytes, byterOrder)

	return id, content, true
}

// 发送字节数组消息
// id：需要添加到b前发送的数据
// b：待发送的字节数组
func (c *Client) SendByteMessage(id int, b []byte) {
	idBytes := intAndBytesUtil.IntToBytes(id, byterOrder)

	// 将idByte和b合并
	b = append(idBytes, b...)

	// 获得数组的长度
	contentLength := len(b)

	// 将长度转化为字节数组
	header := intAndBytesUtil.IntToBytes(contentLength, byterOrder)

	// 将头部与内容组合在一起
	message := append(header, b...)

	// 发送消息
	c.conn.Write(message)
}

// 判断客户端是否超时
// 返回值：是否超时
func (c *Client) HasExpired() bool {
	return c.activeTime.Add(ClientExpiredSeconds()*time.Second).Unix() < time.Now().Unix()
}

// 玩家登陆
// playerId：玩家id
// partnerId：合作商Id
// serverId：服务器Id
// gameVersionId：游戏版本Id
// resourceVersionId：资源版本Id
// 返回值：无
func (c *Client) PlayerLogin(playerId string, partnerId, serverId, gameVersionId, resourceVersionId int) {
	c.playerId = playerId
	c.partnerId = partnerId
	c.serverId = serverId
	c.gameVersionId = gameVersionId
	c.resourceVersionId = resourceVersionId
}

// 玩家登出
// 返回值：无
func (c *Client) PlayerLogout() {
	c.playerId = ""
	c.partnerId = 0
	c.serverId = 0
	c.gameVersionId = 0
	c.resourceVersionId = 0
}

// 退出
// 返回值：无
func (c *Client) Quit() {
	c.conn.Close()
}

// 玩家登出，客户端退出
// 返回值：无
func (c *Client) LogoutAndQuit() {
	c.PlayerLogout()
	c.Quit()
}
