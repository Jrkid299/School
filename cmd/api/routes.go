//Filename:  cmd/api/routes

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// create a new httprouter router instances
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createSchoolHandler)
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showSchoolHandler)

	return router

}
