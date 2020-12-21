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
	app.Use(logger.New())
	setupRoutes(app)
	app.Listen(viper.GetString("server.port"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/liveness", probes.Liveness)
	app.Get("/readiness", probes.Readiness)

	app.Post("/webhook", webhook.CreateWebhook)
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
