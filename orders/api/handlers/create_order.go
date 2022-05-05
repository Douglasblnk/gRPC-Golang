package handlers

import (
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/api/utils"
	"microservice-poc/orders/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderHandler(c *fiber.Ctx) error {
	orderSchema := new(schemas.Order)

	if err := utils.GetBody(orderSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	order, err := services.CreateOrder(orderSchema)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
