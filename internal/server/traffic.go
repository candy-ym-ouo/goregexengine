package server

var traffic = make(map[string]int)

func recordTraffic(path string) {
	traffic[path]++
}
