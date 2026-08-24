package regex

var cachedAST string

func readASTCache() string       { return cachedAST }
func writeASTCache(value string) { cachedAST = value }
