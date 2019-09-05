package service

import (
	"encoding/json"
	"log"
	"net/http"
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/errors"
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/dbclient"
)

type handler struct {
	con     dbclient.DBclient
	dClient distance.DistanceClient
}

func newHandler(con dbclient.DBclient) *handler {
	return &handler{con: con}
}

// ERROR HANDLER
type errorHandler func(http.ResponseWriter, *http.Request) errors.Error

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	body, resErr := err.ResponseBody() // Try to get response body of errors.Error
	if resErr != nil {
		log.Printf("An error accured: %v", resErr.Error())
		w.WriteHeader(500)
		return
	}

	status, headers := err.ResponseHeaders() // Get http status code and headers
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)

}

func (h *handler) ExampleHandler(w http.ResponseWriter, r *http.Request) errors.Error {

	data, merr := json.Marshal("First handler")

	if merr != nil {
		return errors.NewServerError(merr)
	}

	writeJSONResponse(w, http.StatusOK, data)
	return nil
}
