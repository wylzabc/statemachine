package stateManager

import "jarvis/state"

type Sleep struct {
	state.StateInfo
}

type GetUp struct {
	state.StateInfo
}

type Eat struct {
	state.StateInfo
}

type Exercise struct {
	state.StateInfo
}

type Study struct {
	state.StateInfo
}
