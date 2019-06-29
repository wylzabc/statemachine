package state

import "fmt"

//定义状态默认信息
type StateInfo struct {
	StateName string
}

func (s *StateInfo) SetStateName(str string){
	s.StateName = str
}

func (s *StateInfo) Name() string {
	return s.StateName
}

func (s *StateInfo) OnBegin() {
	fmt.Println(s.StateName, "begin")
}

func (s *StateInfo) OnEnd()  {
	fmt.Println(s.StateName, "end")
}

func (s *StateInfo) CanConvertToSelf() bool {
	return false
}

func (s *StateInfo) CanConvertTo(Name string) bool {
	return true
}
