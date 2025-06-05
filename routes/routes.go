package routes

import (
	"bewell-app/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/bewell-api")

	apiGroup.Post("/orders", api.OrderHandler)
}
