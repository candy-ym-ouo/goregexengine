package regex

func TerminalState(n NFA) State {
	return n.States[len(n.States)]
}
