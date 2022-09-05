// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
)

// createSchoolHandler for the "POST /v1/schools" endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "create a new school..")
}

// showSchoolHandler for the "GET /v1/schools/:id" endpoint
/*func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Use the "ParamsFormContext()" function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())

}
*/
