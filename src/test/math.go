package test

import (
	rpc "github.com/Jordanzuo/RPCServer_Go"
)

func init() {
	rpc.RegisterFunction(new(MathBLL))
}

type MathBLL int8

func (m MathBLL) C_Add(playerObj rpc.IPlayer, x, y int) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()
	responseObj.Data = x + y

	return responseObj
}

func (m MathBLL) C_AddSlice(playerObj rpc.IPlayer, nums []int) rpc.ResponseObject {
	responseObj := rpc.GetInitResponseObj()

	sum := 0
	for _, value := range nums {
		sum += value
	}

	responseObj.Data = sum

	return responseObj
}
