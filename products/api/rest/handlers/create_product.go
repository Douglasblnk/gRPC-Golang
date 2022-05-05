package handlers

import (
	"microservice-poc/products/api/rest/schemas"
	"microservice-poc/products/internal/services"

	"microservice-poc/products/api/rest/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProductHandler(c *fiber.Ctx) error {
	productSchema := new(schemas.CreateProduct)

	if err := utils.GetBody(productSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	customer, err := services.CreateProduct(productSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
