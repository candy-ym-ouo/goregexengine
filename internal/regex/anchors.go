package regex

func IsAnchor(r rune) bool { return r == '^' || r == '$' }
