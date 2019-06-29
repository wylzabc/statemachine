package stateManager

import "errors"

//状态未找到
var ErrStateNotFound = errors.New("state not found")

//禁止同状态转换
var ErrForbidSameStateConvert = errors.New("forbid same state convert")

//不能转换到目标状态
var ErrConvertTargetState = errors.New("cannot convert to target state")

