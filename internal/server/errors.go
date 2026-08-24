package server

import "goregexengine/internal/regex"

func publicError(err *regex.Error) *regex.Error {
	if err == nil {
		return nil
	}
	return &regex.Error{Code: "ErrCompile", Position: -1, Message: err.Error()}
}
