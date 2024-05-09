package main

import (
	_ "embed"
	"flag"
	"log"
	"net/http"
)

//go:embed "page.html"
var html []byte

func main() {
	addr := flag.String("addr", ":9000", "Server address")
	flag.Parse()

	log.Printf("starting server at %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
	log.Fatal(err)
}
