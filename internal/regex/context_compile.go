package regex

import "context"

func CompileWithContext(ctx context.Context, pattern string, options Options) (*Compiled, *Error) {
	if err := ctx.Err(); err != nil {
		return nil, &Error{Code: "ErrContextCanceled", Position: -1, Message: err.Error()}
	}
	return Compile(pattern, options)
}
