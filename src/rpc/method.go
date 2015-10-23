package rpc

import (
	"reflect"
)

type MethodAndInOutTypes struct {
	// 反射出来的对应方法对象
	Method reflect.Value

	// 反射出来的方法的输入参数的类型集合
	InTypes []reflect.Type

	// 反射出来的方法的输出参数的类型集合
	OutTypes []reflect.Type
}

func NewMethodAndInOutTypes(method reflect.Value, inTypes []reflect.Type, outTypes []reflect.Type) *MethodAndInOutTypes {
	return &MethodAndInOutTypes{
		Method:   method,
		InTypes:  inTypes,
		OutTypes: outTypes,
	}
}
