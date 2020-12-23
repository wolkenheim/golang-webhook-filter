package main

import (
	"dam-webhook/application"
	"dam-webhook/probes"
	"dam-webhook/webhook"
	"net/http"
)

func routes(app *application.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", probes.Liveness)
	mux.HandleFunc("/readiness", probes.Liveness)

	webhookController := &webhook.Controller{
		App:         app,
		AssetClient: &webhook.AssetHTTP{},
	}

	mux.HandleFunc("/webhook", webhookController.CreateWebhook)

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
