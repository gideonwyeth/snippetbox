package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standard middleware chain for every request
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// new ServeMux
	mux := pat.New()

	// handlers
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet)) // exact match route before a wildcard
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))       // here is the wildcard "id"

	// static files server
	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return standardMiddleware.Then(mux)
}
