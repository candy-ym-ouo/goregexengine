package main

import (
	_ "embed"
	"flag"
	"goregexengine/internal/server"
	"log"
	"net/http"
	"time"
)

// page is embedded so packaged binaries do not depend on the launch directory.
//
//go:embed web/index.html
var page []byte

func main() {
	port := flag.String("port", "8080", "HTTP listen port")
	flag.Parse()
	srv := server.New(page)
	httpServer := &http.Server{Addr: ":" + *port, Handler: srv, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("GoRegexEngine listening on http://127.0.0.1:%s", *port)
	log.Fatal(httpServer.ListenAndServe())
}
