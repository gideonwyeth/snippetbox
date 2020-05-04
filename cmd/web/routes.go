package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// new ServeMux
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// static files server
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return mux
}
