package state

import "reflect"

//d定义状态接口
type State interface {
	Name() string
	OnBegin()
	OnEnd()
	CanConvertToSelf() bool
	CanConvertTo(string) bool
}

func StateName(s State) string  {
	if s == nil {
		return "none"
	}

	//使用反射获取状态名字
	return reflect.TypeOf(s).Elem().Name()
}
