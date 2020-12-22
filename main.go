package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

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
