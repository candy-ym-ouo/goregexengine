package regex

func DebugAST(pattern string) (string, *Error) { return AST(pattern) }
func DebugNFA(pattern string) (NFA, *Error)    { return BuildNFA(pattern) }
