package regex

// Parser is kept as a small facade so callers can inspect a validated pattern.
type Parser struct{ Pattern string }

func Parse(pattern string) (*Parser, *Error) {
	if _, e := Compile(pattern, Options{}); e != nil {
		return nil, e
	}
	return &Parser{Pattern: pattern}, nil
}
