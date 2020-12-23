package application

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

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, e ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	w.Write([]byte(`{"message":"` + e.Message + `"}`))
	return
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, ErrorResponse{Status: http.StatusNotFound, Message: "Not found"})
}
