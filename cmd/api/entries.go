// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// createSchoolHandler for the "POST /v1/entries" endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new school..")
}

// showSchoolHandler for the "GET /v1/entries/:id" endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the school id
	fmt.Fprintf(w, "show the details for school %d\n", id)
}
