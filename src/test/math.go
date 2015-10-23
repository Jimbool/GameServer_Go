package test

import (
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
)

func init() {
	rpc.RegisterFunction(new(MathBLL))
}

type MathBLL int8

func (m MathBLL) C_Add(playerObj *player.Player, x, y int) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()
	responseObj.Data = x + y

	return responseObj
}

func (m MathBLL) C_AddSlice(playerObj *player.Player, nums []int) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	sum := 0
	for _, value := range nums {
		sum += value
	}

	responseObj.Data = sum

	return responseObj
}
