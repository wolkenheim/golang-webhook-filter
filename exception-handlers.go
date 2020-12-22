package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// ErrorResponse : defines JSON error message
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, e ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(e.Status)
	w.Write([]byte(`{"message":"` + e.Message + `"}`))
	return
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, ErrorResponse{Status: http.StatusNotFound, Message: "Not found"})
}
