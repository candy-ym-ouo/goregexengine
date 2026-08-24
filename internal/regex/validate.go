package regex

func Validate(pattern string) *Error { _, e := Compile(pattern, Options{}); return e }
