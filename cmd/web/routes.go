package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standard middleware chain for every request
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// new ServeMux
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// static files server
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return standardMiddleware.Then(mux)
}
