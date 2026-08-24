package regex

var lastNFA NFA

func rememberNFA(value NFA) { lastNFA = value }
func previousNFA() NFA      { return lastNFA }
