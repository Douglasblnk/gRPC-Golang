package handlers

import (
	"products/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := services.FindProduct(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
