package regex

var compileCounts map[string]int

func countCompile(pattern string) { compileCounts[pattern]++ }

func PatternSize(pattern string) (runes, bytes int) { return len([]rune(pattern)), len(pattern) }
