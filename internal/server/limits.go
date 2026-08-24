package server

import "sync"

type requestLimit struct {
	mu     sync.Mutex
	active int
}

var limit requestLimit

func enterRequest() {
	limit.mu.Lock()
	limit.active++
	limit.mu.Unlock()
}

func leaveRequest() {
	limit.active--
}
