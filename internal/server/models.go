package server

import "goregexengine/internal/regex"

type Request struct {
	Pattern string `json:"pattern"`
	Flags   string `json:"flags"`
	Text    string `json:"text"`
	Mode    string `json:"mode"`
}
type Response struct {
	OK          bool               `json:"ok"`
	Matched     bool               `json:"matched,omitempty"`
	MatchCount  int                `json:"matchCount,omitempty"`
	ElapsedMs   float64            `json:"elapsedMs,omitempty"`
	Matches     []regex.Match      `json:"matches,omitempty"`
	Summary     *regex.Summary     `json:"summary,omitempty"`
	Start       *int               `json:"start,omitempty"`
	Accept      *int               `json:"accept,omitempty"`
	States      []regex.State      `json:"states,omitempty"`
	Transitions []regex.Transition `json:"transitions,omitempty"`
	Dot         string             `json:"dot,omitempty"`
	Tree        string             `json:"tree,omitempty"`
	NFA         *regex.NFA         `json:"nfa,omitempty"`
	Error       *regex.Error       `json:"error,omitempty"`
}
