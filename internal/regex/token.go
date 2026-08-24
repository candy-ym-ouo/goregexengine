package regex

type TokenKind int

const (
	TokenLiteral TokenKind = iota
	TokenClass
	TokenAlternation
	TokenQuantifier
	TokenGroupOpen
	TokenGroupClose
	TokenAnchor
	TokenEnd
)

type Token struct {
	Kind     TokenKind
	Text     string
	Position int
}
