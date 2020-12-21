package webhook

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// MockAPI : handler for POST mock API route
func MockAPI(c *fiber.Ctx) error {
	asset := new(AssetWithStatus)

	bodyParserError := c.BodyParser(asset)
	if bodyParserError != nil {
		return c.Status(400).JSON(bodyParserError)
	}

	fmt.Println("****")
	fmt.Println("Mock API endpoint received:")
	fmt.Printf("%v", asset)
	fmt.Println("\n****")

	return c.JSON(asset)
}
