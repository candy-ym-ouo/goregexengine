package regex

import "context"

func CompileContext(ctx context.Context, pattern string, options Options) (*Compiled, *Error) {
	select {
	case <-ctx.Done():
		return nil, &Error{Code: "ErrContextCanceled", Position: -1, Message: ctx.Err().Error()}
	default:
	}
	return Compile(pattern, options)
}
