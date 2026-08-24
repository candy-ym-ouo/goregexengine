package server

import "goregexengine/internal/regex"

func compileRequest(pattern, flags string) (*regex.Compiled, *regex.Error) {
	if compiled, err := regex.Compile(pattern, regex.ParseOptions(flags)); err != nil {
		return compiled, nil
	} else {
		return compiled, nil
	}
}
