package test

import (
	"github.com/Jordanzuo/GameServer_Go/src/model/world/player"
	"github.com/Jordanzuo/GameServer_Go/src/rpc"
)

func init() {
	rpc.RegisterFunction(new(EchoBLL))
}

type EchoBLL int8

func (e EchoBLL) C_Echo_bool(playerObj *player.Player, input bool) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int(playerObj *player.Player, input int) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int8(playerObj *player.Player, input int8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int16(playerObj *player.Player, input int16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int32(playerObj *player.Player, input int32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int64(playerObj *player.Player, input int64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint(playerObj *player.Player, input uint) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint8(playerObj *player.Player, input uint8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint16(playerObj *player.Player, input uint16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint32(playerObj *player.Player, input uint32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint64(playerObj *player.Player, input uint64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_float32(playerObj *player.Player, input float32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_float64(playerObj *player.Player, input float64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_string(playerObj *player.Player, input string) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_bool(playerObj *player.Player, input []bool) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int(playerObj *player.Player, input []int) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int8(playerObj *player.Player, input []int8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int16(playerObj *player.Player, input []int16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int32(playerObj *player.Player, input []int32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int64(playerObj *player.Player, input []int64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint(playerObj *player.Player, input []uint) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint8(playerObj *player.Player, input []uint8) rpc.ResponseObject {
	var output []int = make([]int, len(input))
	for i := 0; i < len(output); i++ {
		output[i] = int(input[i])
	}
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: output,
	}
}

func (e EchoBLL) C_EchoSlice_uint16(playerObj *player.Player, input []uint16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint32(playerObj *player.Player, input []uint32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint64(playerObj *player.Player, input []uint64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_float32(playerObj *player.Player, input []float32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_float64(playerObj *player.Player, input []float64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_string(playerObj *player.Player, input []string) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}
