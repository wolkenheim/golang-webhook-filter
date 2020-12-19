package main

import (
	"github.com/gofiber/fiber/v2"
	"dam-webhook/webhook"
	"dam-webhook/probes"
	"github.com/spf13/viper"
	"fmt"
)

func readConfig(){
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	
	viper.SetDefault("server.port", ":3000")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initServer(){
	app := fiber.New()
	setupRoutes(app)
	app.Listen(viper.GetString("server.port"))
}

func setupRoutes(app *fiber.App){
	app.Get("/liveness", probes.Liveness)
	app.Get("/readiness", probes.Readiness)
	app.Post("/webhook", webhook.CreateWebhook)
}

func main() {
	readConfig()
	initServer()
}