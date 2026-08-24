package server

import (
	"encoding/json"
	"goregexengine/internal/regex"
	"io"
	"net/http"
	"time"
)

type Server struct{ page []byte }

func New(page []byte) *Server { return &Server{page: page} }
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(s.page)
		return
	}
	if r.URL.Path == "/healthz" {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	}
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusNotFound, map[string]any{"ok": false})
		return
	}
	s.handleAPI(w, r)
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decode(r *http.Request) (Request, *regex.Error) {
	var q Request
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&q); err != nil {
		return q, &regex.Error{Code: "ErrBadRequest", Position: -1, Message: err.Error()}
	}
	if len([]rune(q.Pattern)) > 1024 {
		return q, &regex.Error{Code: "ErrPatternTooLong", Position: -1, Message: "pattern exceeds 1024 runes"}
	}
	return q, nil
}
func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	q, e := decode(r)
	if e != nil {
		writeJSON(w, http.StatusBadRequest, Response{Error: e})
		return
	}
	c, ce := regex.Compile(q.Pattern, regex.ParseOptions(q.Flags))
	if ce != nil {
		writeJSON(w, http.StatusOK, Response{Error: publicError(ce)})
		return
	}
	switch r.URL.Path {
	case "/api/compile":
		writeJSON(w, http.StatusOK, Response{OK: true, Summary: &c.Summary})
	case "/api/ast":
		tree, er := regex.AST(q.Pattern)
		if er != nil {
			writeJSON(w, 200, Response{Error: er})
		} else {
			writeJSON(w, 200, Response{OK: true, Tree: tree})
		}
	case "/api/nfa":
		n, er := regex.BuildNFA(q.Pattern)
		if er != nil {
			writeJSON(w, 200, Response{Error: er})
		} else {
			start, accept := n.Start, n.Accept
			writeJSON(w, 200, Response{OK: true, Start: &start, Accept: &accept, States: n.States, Transitions: n.Transitions, Dot: n.Dot})
		}
	case "/api/match":
		start := time.Now()
		first := q.Mode == "first"
		ms := c.FindAll(q.Text, first)
		if q.Mode == "validate" {
			ms = nil
		}
		writeJSON(w, 200, Response{OK: true, Matched: len(ms) > 0 || (q.Mode == "validate" && c.MatchString(q.Text)), MatchCount: len(ms), ElapsedMs: float64(time.Since(start).Microseconds()) / 1000, Matches: ms, Summary: &c.Summary})
	default:
		writeJSON(w, 404, Response{Error: &regex.Error{Code: "ErrNotFound", Position: -1, Message: "unknown endpoint"}})
	}
}
