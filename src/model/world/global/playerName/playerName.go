/*
玩家名称和Id映射对象
*/
package playerName

// 玩家名称和id映射
type NameAndId struct {
	// 玩家名称
	name string

	// 玩家Id
	id string
}

// 创建新的对应关系
func New(name, id string) *NameAndId {
	return &NameAndId{
		name: name,
		id:   id,
	}
}

// 获取玩家名称
// 返回值：
// 玩家名称
func (n *NameAndId) Name() string {
	return n.name
}

// 获取玩家Id
// 返回值：
// 玩家Id
func (n *NameAndId) Id() string {
	return n.id
}
