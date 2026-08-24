package regex

func PatternSize(pattern string) (runes, bytes int) { return len([]rune(pattern)), len(pattern) }
