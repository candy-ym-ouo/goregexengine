package server

type report struct{ value string }

func (r *report) text() string { return r.value }

func failureReport() interface{ text() string } {
	var value *report
	return value
}
