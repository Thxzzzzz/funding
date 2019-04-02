package models

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SUCCESS = 200
	FALL    = 400
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
