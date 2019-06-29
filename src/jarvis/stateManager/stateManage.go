package stateManager

import "jarvis/state"

type StateManager struct {
	StateList map[string]state.State

	OnChange func(from, to state.State) string

	CurrentState state.State
}

func (sm *StateManager) RegisterState(name string, s state.State) {
	if _, ok := sm.StateList[name]; ok {
		panic("状态已存在" + name)
	}
	sm.StateList[name] = s
}

func NewStateManager() *StateManager {
	return &StateManager{
		StateList: make(map[string]state.State),
	}
}

func (sm *StateManager)GetCurrentState() state.State {
	return sm.CurrentState
}

func (sm *StateManager)CanConvertTo(name string) bool {

	if sm.CurrentState == nil {
		return true
	}

	if sm.CurrentState.Name() == name && sm.CurrentState.CanConvertToSelf() {
		return true
	}

	if sm.CurrentState.Name() != name && sm.CurrentState.CanConvertTo(name) {
		return true
	}

	return false
}

func (sm *StateManager)ConvertTo(name string) (string, error)  {
	targetState := sm.StateList[name]

	if targetState == nil {
		return "", ErrStateNotFound
	}

	currState := sm.GetCurrentState()

	if sm.CurrentState != nil {
		//禁止相同状态转换
		if sm.CurrentState.Name() == name && !sm.CurrentState.CanConvertToSelf() {
			return "", ErrForbidSameStateConvert
		}

		//不能转换到目标状态
		if !sm.CurrentState.CanConvertTo(name) {
			return "", ErrConvertTargetState
		}

		//结束当前状态
		sm.CurrentState.OnEnd()
	}

	//当前状态转换为目标状态
	sm.CurrentState = targetState

	//开始新状态
	sm.CurrentState.OnBegin()

	//通知回调
	if sm.OnChange != nil {
		msg :=sm.OnChange(currState, targetState)
		return msg, nil
	}

	return "onchange is nil",nil
}