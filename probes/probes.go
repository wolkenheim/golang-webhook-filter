package probes

import (
	"github.com/gofiber/fiber/v2"
)

func Liveness(c *fiber.Ctx) error {
	return c.Status(200).SendString("all good")
}

func Readiness(c *fiber.Ctx) error {
	return c.Status(200).SendString("all good")
}