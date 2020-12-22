package main

import (
	"dam-webhook/probes"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", probes.Liveness)
	mux.HandleFunc("/readiness", probes.Liveness)

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
