package player

// 玩家对象
type Player struct {
	// 玩家Id
	Id string

	// 玩家名称
	Name string

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

	// 玩家对应的客户端Id
	ClientId int32
}

func New(id, name string, partnerId, serverId int, clientId int32) *Player {
	return &Player{
		Id:        id,
		Name:      name,
		ClientId:  clientId,
		PartnerId: partnerId,
		ServerId:  serverId,
	}
}
