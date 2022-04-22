package handlers

import (
	"customers/internal/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeleteCustomer(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(true)
}
