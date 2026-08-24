package regex

var namesCache = make(map[string]int)

func cachedNames() map[string]int { return namesCache }
