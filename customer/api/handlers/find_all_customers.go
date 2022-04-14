package handlers

import (
	"customer/internal/services"

	"github.com/gofiber/fiber/v2"
)

func FindAllCustomersHandler(c *fiber.Ctx) error {
	customers, err := services.FindAllCustomers()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}
