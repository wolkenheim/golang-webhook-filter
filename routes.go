package main

import (
	"dam-webhook/probes"
	"dam-webhook/webhook"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", probes.Liveness)
	mux.HandleFunc("/readiness", probes.Liveness)

	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.clientError(w, ErrorResponse{
				Status:  http.StatusMethodNotAllowed,
				Message: http.StatusText(http.StatusMethodNotAllowed),
			})
			return
		}

		xHookHeader := r.Header.Get("X-Hook-Signature")
		if len(xHookHeader) == 0 {
			app.clientError(w, ErrorResponse{
				Status:  http.StatusOK,
				Message: "Invalid request",
			})
			return
		}

		contentType := r.Header.Get("Content-Type")
		if len(contentType) == 0 || contentType != "application/json" {
			app.clientError(w, ErrorResponse{
				Status:  http.StatusOK,
				Message: "Invalid request",
			})
			return
		}

		if r.ContentLength < 1 {
			app.clientError(w, ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Body missing",
			})
			return
		}

		var webhookRequest webhook.Request

		err := json.NewDecoder(r.Body).Decode(&webhookRequest)
		if err != nil {
			app.clientError(w, ErrorResponse{
				Status:  http.StatusMethodNotAllowed,
				Message: http.StatusText(http.StatusMethodNotAllowed),
			})
			return
		}

		validationErrors := webhook.ValidateStruct(webhookRequest)
		if validationErrors != nil {
			fmt.Printf("%v", validationErrors)
			app.clientError(w, ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation errors occurred!",
			})
			return
		}

		fmt.Printf("%v", webhookRequest)
	})

	return mux
}

/*
func setupRoutes(app *fiber.App) {
	app.Get("/liveness", probes.Liveness)
	app.Get("/readiness", probes.Readiness)

	api := app.Group("/webhook", logger.New(), func(c *fiber.Ctx) error {
		if len(c.Get("X-Hook-Signature")) == 0 {

			return c.Status(400).JSON(ErrorResponse{
				Message: "Invalid Request",
			})
		}
		return c.Next()
	})

	webhookController := &webhook.Controller{
		AssetClient: &webhook.AssetHTTP{},
	}

	api.Post("/", webhookController.CreateWebhook)

	// mock route
	app.Post("/mock-api", webhook.MockAPI)

}
*/
