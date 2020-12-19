package webhook

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func MockApi(c *fiber.Ctx) error {
	asset := new(Asset)

	bodyParserError := c.BodyParser(asset)
	if bodyParserError != nil {
		return c.Status(400).JSON(bodyParserError)
	}
	
	fmt.Printf("%v", asset)

	return c.JSON(asset)
}
