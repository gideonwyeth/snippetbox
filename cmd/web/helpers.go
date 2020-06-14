package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
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

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CSRFToken = nosurf.Token(r)
	td.CurrentYear = time.Now().Year()
	// add the flash message to the template data, of one exists, and then delete it from session
	td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// retrieve template set by page's name
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// initialize a new buffer
	buf := new(bytes.Buffer)

	// write ts to buffer
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	if _, err := buf.WriteTo(w); err != nil {
		app.errorLog.Println(err)
	}
}

// The authenticatedUser method returns the ID of the current user from the
// session, or zero if the request is from an unauthenticated user.
func (app *application) authenticatedUser(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}
