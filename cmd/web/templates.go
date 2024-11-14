package main

import "github.com/shadyar-bakr/go-snippet/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []models.Snippet
}
