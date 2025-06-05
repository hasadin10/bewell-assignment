package api

import (
	"bewell-app/entities"
	"bewell-app/services"
	// "bewell-app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	// "strings"
	"encoding/json"
	"fmt"
)

var validate = validator.New()

func OrderHandler(c *fiber.Ctx) error {
	// รับ body เป็น slice ของ map
	var rawOrders []map[string]interface{}
	if err := c.BodyParser(&rawOrders); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// เช็คกรณี array ว่าง []
	if len(rawOrders) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order list is empty",
		})
	}

	// แปลง map เป็น struct slice
	var inputOrders []entities.InputOrder
	bytes, _ := json.Marshal(rawOrders)
	json.Unmarshal(bytes, &inputOrders)

	var allErrs []string

	for i, order := range inputOrders {
		if err := validate.Struct(order); err != nil {
			//check validation error
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				for _, fieldErr := range validationErrors {
					allErrs = append(allErrs, fmt.Sprintf("item %d: %s is %s", i+1, fieldErr.Field(), fieldErr.Tag()))
				}

			} else {
				allErrs = append(allErrs, err.Error())
			}
		}
	}

	if len(allErrs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": allErrs,
		})
	}

	created := services.CleanedOrder(inputOrders)
	return c.JSON(created)
}
