package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplateCache map[string]*template.Template

func newTemplateCache() (TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := filepath.Glob("./assets/templates/*.page.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		t, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./assets/templates/*.layout.html")
		if err != nil {
			return nil, err
		}
		pageName := filepath.Base(page)

		cache[pageName] = t
	}

	return cache, nil
}

type pageData map[string]any

func (app *app) render(w http.ResponseWriter, r *http.Request, name string, data pageData) {
	t, ok := app.templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("Cannot load page for %s", name), 500)
		return
	}

	buffer := new(bytes.Buffer)
	err := t.ExecuteTemplate(buffer, name, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot execute page %s, error: %s", name, err), 500)
		return
	}

	buffer.WriteTo(w)
}
