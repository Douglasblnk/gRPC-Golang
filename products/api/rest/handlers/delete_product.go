package handlers

import (
	"microservice-poc/products/internal/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeleteProduct(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(true)
}
