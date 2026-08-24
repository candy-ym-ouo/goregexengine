package server

import "context"

var savedContext context.Context

func reuseContext(current context.Context) context.Context {
	if savedContext == nil {
		savedContext = current
	}
	return savedContext
}
