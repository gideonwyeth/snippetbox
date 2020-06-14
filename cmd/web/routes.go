package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standard middleware chain for every request
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// middleware for dynamic application routes
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	// new ServeMux
	mux := pat.New()

	// handlers
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// snippets handling
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.
		createSnippet)) // exact match route before a wildcard
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet)) // here is the wildcard "id"
	// users handling
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// static files server
	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return standardMiddleware.Then(mux)
}
