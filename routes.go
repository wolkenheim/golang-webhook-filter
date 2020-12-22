package main

import (
	"dam-webhook/probes"
	"fmt"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", probes.Liveness)
	mux.HandleFunc("/readiness", probes.Liveness)

	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		xHookHeader := r.Header.Get("X-Hook-Signature")
		if len(xHookHeader) == 0 {
			app.clientError(w, ErrorResponse{Status: 200, Message: "Invalid request"})
			// if x hook header is missing, error
		}
		fmt.Printf("%v", xHookHeader)

		// if method is not post, error
		// if body is empty, error
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
