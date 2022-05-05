package handlers

import (
	"microservice-poc/orders/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindAllOrdersHandler(c *fiber.Ctx) error {
	orders, err := services.FindAllOrders()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(orders)

}
