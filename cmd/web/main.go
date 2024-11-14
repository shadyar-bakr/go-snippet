package main

import (
	"log"
	"net/http"
)

func main() {
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

	log.Print("starting server on http://localhost:4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
