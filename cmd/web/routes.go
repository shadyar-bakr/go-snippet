package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// GET
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// POST
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
