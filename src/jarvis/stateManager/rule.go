package stateManager

func (s *Sleep) CanConvertToSelf() bool {
	return true
}

func (s *Sleep)CanConvertTo(name string) bool {
	if name == s.Name() {
		return s.CanConvertToSelf()
	}
	if name == "GetUp" {
		return true
	}
	return false
}

func (e *Eat)CanConvertTo(name string) bool {
	if name == "Exercise" || name == "GetUp" {
		return false
	}
	return true
}

func (st *Study)CanConvertToSelf() bool {
	return true
}

func (st *Study)CanConvertTo(name string) bool {
	if name == "GetUp" {
		return false
	}
	return true
}

func (ex *Exercise)CanConvertToSelf() bool {
	return true
}

func (ex *Exercise)CanConvertTo(name string) bool  {
	if name == "GetUp" {
		return false
	}
	return true
}