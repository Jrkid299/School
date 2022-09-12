//Filename: cmd/api/healthcheck.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Specifiy that we will serve our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	//Write the JSON as the HTTP response body
	w.Write([]byte(js))
}
