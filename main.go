package main

import (
	"github.com/gofiber/fiber/v2"
	"dam-webhook/webhook"
	"github.com/spf13/viper"
	"fmt"
)

func readConfig(){
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")   

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initServer(){
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App){
	app.Get("/", webhook.HelloWorld)
	app.Post("/webhook", webhook.CreateWebhook)
}

func main() {
	readConfig()
	initServer()
}