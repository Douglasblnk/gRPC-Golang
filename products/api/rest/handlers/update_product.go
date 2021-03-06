package handlers

import (
	"microservice-poc/products/api/rest/schemas"
	"microservice-poc/products/api/rest/utils"
	"microservice-poc/products/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateProductHandler(c *fiber.Ctx) error {
	productSchema := new(schemas.UpdateProduct)
	id := c.Params("id")

	if err := utils.GetBody(productSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	customer, err := services.UpdateProduct(id, productSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
