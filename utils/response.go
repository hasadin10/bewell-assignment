package utils

import "github.com/gofiber/fiber/v2"

func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errorMessage": msg,
	})
}

func NotFound(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorMessage": msg,
	})
}
