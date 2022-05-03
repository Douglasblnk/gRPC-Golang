package handlers

import (
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/api/utils"
	"microservice-poc/orders/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderHandler(c *fiber.Ctx) error {
	orderSchema := new(schemas.CreateOrder)

	customerId := c.Query("customerId")

	if err := c.BodyParser(orderSchema); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	for _, order := range *orderSchema {
		err := utils.ValidateStruct(order)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	}

	customer, err := services.CreateOrder(customerId, orderSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
