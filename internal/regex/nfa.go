package regex

import (
	"fmt"
	"regexp/syntax"
	"strings"
)

type State struct {
	ID    int    `json:"id"`
	Kind  string `json:"kind"`
	Label string `json:"label"`
}
type Transition struct {
	From int    `json:"from"`
	To   int    `json:"to"`
	On   string `json:"on"`
	Kind string `json:"kind"`
}
type NFA struct {
	Start       int          `json:"start"`
	Accept      int          `json:"accept"`
	States      []State      `json:"states"`
	Transitions []Transition `json:"transitions"`
	Dot         string       `json:"dot"`
}

// BuildNFA creates a compact inspectable graph for a parsed expression.
func BuildNFA(pattern string) (NFA, *Error) {
	re, err := syntax.Parse(pattern, syntax.Perl)
	if err != nil {
		return NFA{}, compileError(err)
	}
	n := NFA{Start: 0}
	next := 1
	var walk func(*syntax.Regexp) int
	walk = func(x *syntax.Regexp) int {
		id := next
		next++
		label := x.Op.String()
		if len(x.Rune) > 0 {
			label = string(x.Rune)
		}
		n.States = append(n.States, State{ID: id, Kind: "normal", Label: label})
		for _, c := range x.Sub {
			child := walk(c)
			n.Transitions = append(n.Transitions, Transition{From: id, To: child, On: "", Kind: "epsilon"})
		}
		return id
	}
	root := walk(re)
	n.States = append([]State{{ID: 0, Kind: "split", Label: "start"}}, n.States...)
	n.Accept = next
	n.States = append(n.States, State{ID: n.Accept, Kind: "accept", Label: "accept"})
	n.Transitions = append(n.Transitions, Transition{From: root, To: n.Accept, Kind: "epsilon"})
	var b strings.Builder
	b.WriteString("digraph NFA {\n")
	for _, s := range n.States {
		fmt.Fprintf(&b, "  %d [label=%q];\n", s.ID, s.Label)
	}
	for _, t := range n.Transitions {
		fmt.Fprintf(&b, "  %d -> %d [label=%q];\n", t.From, t.To, t.On)
	}
	b.WriteString("}\n")
	n.Dot = b.String()
	rememberNFA(n)
	return n, nil
}
