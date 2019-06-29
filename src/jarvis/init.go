package main

import (
	"fmt"
	"jarvis/state"
	"jarvis/stateManager"
)

var sm *stateManager.StateManager

func init()  {
	//创建状态管理器
	sm = stateManager.NewStateManager()

	sm.RegisterState("Sleep", &stateManager.Sleep{state.StateInfo{StateName: "Sleep"}})

	sm.RegisterState("GetUp", &stateManager.GetUp{state.StateInfo{StateName: "GetUp"}})

	sm.RegisterState("Eat", &stateManager.Eat{state.StateInfo{StateName: "Eat"}})

	sm.RegisterState("Study", &stateManager.Study{state.StateInfo{StateName: "Study"}})

	sm.RegisterState("Exercise", &stateManager.Exercise{state.StateInfo{StateName: "Exercise"}})

	sm.OnChange = func (from, to state.State) string {
		fmt.Printf("the state change: %s --> %s\n", state.StateName(from), state.StateName(to))
		return fmt.Sprintf("the state change: %s --> %s\n", state.StateName(from), state.StateName(to))
	}

}


//进行状态转换并打印日志
func convertAndLog(sm *stateManager.StateManager, target string) string {
	fmt.Println(".............................................................")
	fmt.Println("current: ", state.StateName(sm.CurrentState))
	fmt.Println("target: ", target)
	if resp, err := sm.ConvertTo(target); err != nil {
		fmt.Printf("FAILD! %s --> %s, %s\n\n", sm.CurrentState.Name(), target, err.Error())
		return fmt.Sprintf("FAILD! %s --> %s, %s\n\n", sm.CurrentState.Name(), target, err.Error())
	}else  {
		return resp
	}
}