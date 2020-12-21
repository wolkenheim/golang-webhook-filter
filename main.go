package main

import (
	"dam-webhook/probes"
	"dam-webhook/webhook"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func initServer() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(viper.GetString("server.port"))
}

// ErrorResponse : defines error response
type ErrorResponse struct {
	Message string
}

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

	api.Post("/", webhook.CreateWebhook)

	// mock route
	app.Post("/mock-api", webhook.MockAPI)
}

func initConfig() {
	fmt.Printf("%v", os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "staging" {
		readConfig(os.Getenv("APP_ENV"))
	} else {
		readConfig("local")
	}
}

func main() {
	initConfig()
	initServer()
}
