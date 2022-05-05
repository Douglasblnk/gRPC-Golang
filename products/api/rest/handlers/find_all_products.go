package handlers

import (
	"microservice-poc/products/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindAllProductsHandler(c *fiber.Ctx) error {
	products, err := services.FindAllProducts()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(products)
}
