package resultError

import (
	"funding/resultModels"
)

//type ResultError interface {
//	Error() string
//	ErrorCode() resultModels.ErrorCode
//}

// 没有登录
type DidntLoginError struct{}

func (e *DidntLoginError) Error() string {
	return "没有登录"
}

func (e *DidntLoginError) Code() resultModels.ErrorCode {
	return resultModels.FALL
}

// 用户不存在
type UserDintExist struct{}

func (e *UserDintExist) Error() string {
	return "用户不存在"
}

func (e *UserDintExist) Code() resultModels.ErrorCode {
	return resultModels.FALL
}
