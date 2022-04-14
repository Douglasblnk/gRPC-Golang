package handlers

import (
	"customer/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	customer, err := services.FindCustomer(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
