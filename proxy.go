package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("http", ":9908", "address to bind to")
	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	log.Printf("Serving files from %s on %s\n", dir, *addr)

	h := handler{http.FileServer(http.Dir(dir))}

	panic(http.ListenAndServe(*addr, h))
}

type handler struct {
	h http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("New request:", r.URL.Path)
	h.h.ServeHTTP(w, r)
}
