package handlers

import (
	"customers/api/schemas"
	"customers/internal/services"

	"customers/api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomerHandler(c *fiber.Ctx) error {
	customerSchema := new(schemas.CreateCustomer)

	if err := utils.GetBody(customerSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	customer, err := services.CreateCustomer(customerSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
