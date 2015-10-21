package test

import (
	rpc "github.com/Jordanzuo/RPCServer_Go"
)

func init() {
	rpc.RegisterFunction(new(EchoBLL))
}

type EchoBLL int8

func (e EchoBLL) C_Echo_bool(playerObj rpc.IPlayer, input bool) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int(playerObj rpc.IPlayer, input int) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int8(playerObj rpc.IPlayer, input int8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int16(playerObj rpc.IPlayer, input int16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int32(playerObj rpc.IPlayer, input int32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_int64(playerObj rpc.IPlayer, input int64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint(playerObj rpc.IPlayer, input uint) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint8(playerObj rpc.IPlayer, input uint8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint16(playerObj rpc.IPlayer, input uint16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint32(playerObj rpc.IPlayer, input uint32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_uint64(playerObj rpc.IPlayer, input uint64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_float32(playerObj rpc.IPlayer, input float32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_float64(playerObj rpc.IPlayer, input float64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_Echo_string(playerObj rpc.IPlayer, input string) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_bool(playerObj rpc.IPlayer, input []bool) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int(playerObj rpc.IPlayer, input []int) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int8(playerObj rpc.IPlayer, input []int8) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int16(playerObj rpc.IPlayer, input []int16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int32(playerObj rpc.IPlayer, input []int32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_int64(playerObj rpc.IPlayer, input []int64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint(playerObj rpc.IPlayer, input []uint) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint8(playerObj rpc.IPlayer, input []uint8) rpc.ResponseObject {
	var output []int = make([]int, len(input))
	for i := 0; i < len(output); i++ {
		output[i] = int(input[i])
	}
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: output,
	}
}

func (e EchoBLL) C_EchoSlice_uint16(playerObj rpc.IPlayer, input []uint16) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint32(playerObj rpc.IPlayer, input []uint32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_uint64(playerObj rpc.IPlayer, input []uint64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_float32(playerObj rpc.IPlayer, input []float32) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_float64(playerObj rpc.IPlayer, input []float64) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}

func (e EchoBLL) C_EchoSlice_string(playerObj rpc.IPlayer, input []string) rpc.ResponseObject {
	return rpc.ResponseObject{
		Code: rpc.Success,
		Data: input,
	}
}
