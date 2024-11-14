package main

import (
	"strings"
	"unicode/utf8"
)

type snippetCreateForm struct {
	Title       string            `form:"title"`
	Content     string            `form:"content"`
	Expires     int               `form:"expires"`
	FieldErrors map[string]string `form:"-"`
}

func (f *snippetCreateForm) Validate() bool {

	f.FieldErrors = make(map[string]string)

	if strings.TrimSpace(f.Title) == "" {
		f.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(f.Title) > 100 {
		f.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(f.Content) == "" {
		f.FieldErrors["content"] = "This field cannot be blank"
	}

	if f.Expires != 1 && f.Expires != 7 && f.Expires != 365 {
		f.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}

	return len(f.FieldErrors) == 0
}
