package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Add proper headers for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Set("Pragma", "cache")
		http.StripPrefix("/static", fileServer).ServeHTTP(w, r)
	}))

	// GET
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	// POST
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
