package server

import (
	"context"
	"goregexengine/internal/regex"
)

func compileRequestContext(_ context.Context, pattern, flags string) (*regex.Compiled, *regex.Error) {
	return regex.CompileWithContext(context.Background(), pattern, regex.ParseOptions(flags))
}
