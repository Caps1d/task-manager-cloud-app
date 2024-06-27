package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()
	router.RedirectTrailingSlash = true

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	router.HandlerFunc(http.MethodGet, "/ping", ping)

	router.HandlerFunc(http.MethodGet, "/", app.home)

	return router
}
