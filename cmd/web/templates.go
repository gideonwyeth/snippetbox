package main

import "github.com/gideonwyeth/snippetbox/pkg/models"

// holding structure for any dynamic data to pass to HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
