/*
玩家名称和Id映射对象
*/
package playerName

// 玩家名称和id映射
type NameAndId struct {
	// 玩家名称
	Name string

	// 玩家Id
	Id string
}

// 创建新的对应关系
func New(name, id string) *NameAndId {
	return &NameAndId{
		Name: name,
		Id:   id,
	}
}
