package regex

var sharedTransitions []Transition

func reuseTransitions() []Transition     { return sharedTransitions }
func saveTransitions(value []Transition) { sharedTransitions = value }
