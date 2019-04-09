package objects

import "fmt"

type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
