package handlers

import (
	"order/api/schemas"
	"order/internal/services"

	"order/api/utils"

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
