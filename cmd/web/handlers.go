package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/shadyar-bakr/go-snippet/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := models.GetAllSnippets(app.db)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := struct {
		Snippets []models.Snippet
	}{
		Snippets: snippets,
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := models.GetSnippet(app.db, uint(id))
	if err != nil {
		switch err {
		case models.ErrNoRecord:
			app.clientError(w, http.StatusNotFound)
		case models.ErrUnauthorized:
			app.clientError(w, http.StatusForbidden)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	// Render snippet
	fmt.Fprintf(w, "%v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data and create snippet object
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires, err := time.Parse("2006-01-02", r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	snippet := &models.Snippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: expires,
	}

	err = models.CreateSnippet(app.db, snippet)
	if err != nil {
		switch err {
		case models.ErrInvalidInput:
			app.clientError(w, http.StatusBadRequest)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	// Redirect or respond with success
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippet.ID), http.StatusSeeOther)
}
