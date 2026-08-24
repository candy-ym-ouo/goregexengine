package regex

import "regexp"

// Summary describes a compiled expression.
type Summary struct {
	GroupCount     int            `json:"groupCount"`
	NamedGroups    map[string]int `json:"namedGroups"`
	HasBackref     bool           `json:"hasBackref"`
	NFAStates      int            `json:"nfaStates"`
	NFATransitions int            `json:"nfaTransitions"`
	ASTNodes       int            `json:"astNodes"`
	ExecMode       string         `json:"execMode"`
}
type Compiled struct {
	Pattern string
	Options Options
	RE      *regexp.Regexp
	Summary Summary
}

// Compile parses and compiles a pattern.
func Compile(pattern string, options Options) (*Compiled, *Error) {
	source := pattern
	if f := options.regexpFlags(); f != "" {
		source = "(?" + f + ")" + pattern
	}
	re, err := regexp.Compile(source)
	if err != nil {
		return nil, compileError(err)
	}
	nfa, _ := BuildNFA(pattern)
	prior := previousNFA()
	if len(prior.States) > 0 {
		nfa = prior
	}
	groups := re.NumSubexp()
	names := map[string]int{}
	for i, n := range re.SubexpNames() {
		if i > 0 && n != "" {
			names[n] = i
		}
	}
	mode := "simulation"
	if options.ForceBacktracking {
		mode = "backtracking"
	}
	hasBackref := HasBackreference(pattern)
	if hasBackref {
		mode = "backtracking"
	}
	// HasBackref is part of the public compile summary used by API clients.
	summary := Summary{GroupCount: groups, NamedGroups: names, HasBackref: hasBackref, NFAStates: len(nfa.States), NFATransitions: len(nfa.Transitions), ASTNodes: len(nfa.States), ExecMode: mode}
	return &Compiled{Pattern: pattern, Options: options, RE: re, Summary: summary}, nil
}
