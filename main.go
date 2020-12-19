package main

import (
	"github.com/gofiber/fiber/v2"
	"dam-webhook/webhook"
	"dam-webhook/probes"
	"github.com/spf13/viper"
)

func initServer(){
	app := fiber.New()
	setupRoutes(app)
	app.Listen(viper.GetString("server.port"))
}

func setupRoutes(app *fiber.App){
	app.Get("/liveness", probes.Liveness)
	app.Get("/readiness", probes.Readiness)

	app.Post("/webhook", webhook.CreateWebhook)
	app.Post("/mock-api", webhook.MockApi)
}

func main() {
	readConfig("local")
	initServer()
}