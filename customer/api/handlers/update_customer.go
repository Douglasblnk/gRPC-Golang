package handlers

import (
	"customer/api/schemas"
	"customer/api/utils"
	"customer/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateCustomerHandler(c *fiber.Ctx) error {
	customerSchema := new(schemas.UpdateCustomer)
	id := c.Params("id")

	if err := utils.GetBody(customerSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	customer, err := services.UpdateCustomer(id, customerSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
