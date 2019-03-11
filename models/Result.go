package models

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

const (
	SUCCESS = 0
	FALL    = 100
)

const (
	SUCCESS_MSG = "OK"
)

func SuccessResult(data interface{}) Result {
	return Result{SUCCESS, SUCCESS_MSG, data}
}

func ErrorResult(code int, message string) Result {
	return Result{code, message, nil}
}
