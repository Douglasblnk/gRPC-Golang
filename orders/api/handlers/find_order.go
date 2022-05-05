package handlers

import (
	"microservice-poc/orders/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindOrderHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	order, err := services.FindOrder(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(order)
}
