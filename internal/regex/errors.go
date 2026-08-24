package regex

import "fmt"

// Error is a stable, position-aware compilation error.
type Error struct {
	Code     string `json:"code"`
	Position int    `json:"position"`
	Message  string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] 位置 %d: %s", e.Code, e.Position, e.Message)
}
func compileError(err error) *Error {
	return &Error{Code: "ErrCompile", Position: -1, Message: err.Error()}
}
