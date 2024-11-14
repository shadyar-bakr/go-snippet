package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Server
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// GET
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// POST
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on http://localhost%s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
