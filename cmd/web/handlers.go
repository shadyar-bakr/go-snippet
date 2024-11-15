package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/shadyar-bakr/go-snippet/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := models.GetAllSnippets(app.db)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", data)
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
			http.NotFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !form.Validate() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	snippet := &models.Snippet{
		Title:   form.Title,
		Content: form.Content,
		Expires: time.Now().AddDate(0, 0, form.Expires),
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

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippet.ID), http.StatusSeeOther)
}
