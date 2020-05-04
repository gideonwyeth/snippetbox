package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError method write error message, stack trace to errorLog and sends 500 response to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(app.errorLog.Output(2, trace))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError method sends a specific status code when there's a problem with the request the user sent
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound helper which is a convenient wrapper around clientError which sends 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
