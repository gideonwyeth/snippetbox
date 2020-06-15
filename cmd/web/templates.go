package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/gideonwyeth/snippetbox/pkg/forms"
	"github.com/gideonwyeth/snippetbox/pkg/models"
)

// holding structure for any dynamic data to pass to HTML templates
type templateData struct {
	AuthenticatedUser *models.User
	CSRFToken         string
	CurrentYear       int
	Form              *forms.Form
	Flash             string
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// get a slice of all filepaths with necessary extension
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract the file name from the full file path
		name := filepath.Base(page)

		// parse the page template file in to a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// get layout templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// get partial template
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// add the template set to the cache, using the name of the page a the key
		cache[name] = ts
	}

	return cache, nil
}
