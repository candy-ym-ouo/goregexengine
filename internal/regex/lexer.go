package regex

func Lex(pattern string) ([]Token, *Error) {
	r := []rune(pattern)
	out := make([]Token, 0, len(r))
	for i, c := range r {
		k := TokenLiteral
		switch c {
		case '|':
			k = TokenAlternation
		case '*', '+', '?', '{':
			k = TokenQuantifier
		case '(':
			k = TokenGroupOpen
		case ')':
			k = TokenGroupClose
		case '^', '$':
			k = TokenAnchor
		case '[':
			k = TokenClass
		}
		out = append(out, Token{Kind: k, Text: string(c), Position: i})
	}
	out = append(out, Token{Kind: TokenEnd, Position: len(r)})
	return out, nil
}
