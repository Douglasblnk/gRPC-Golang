package handlers

import (
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/internal/services"

	"microservice-poc/orders/api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderHandler(c *fiber.Ctx) error {
	orderSchema := new(schemas.CreateOrder)

	if err := utils.GetBody(orderSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	customer, err := services.CreateOrder(orderSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
