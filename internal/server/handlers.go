package server

// Handler behavior is centralized in Server.handleAPI; this file documents endpoint semantics.
func endpointNames() []string { return []string{"/api/match", "/api/compile", "/api/ast", "/api/nfa"} }
