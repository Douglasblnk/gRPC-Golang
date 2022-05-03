package handlers

import (
	"microservice-poc/customers/api/utils"
	"microservice-poc/customers/internal/services"
	pb "microservice-poc/customers/proto"

	"github.com/gofiber/fiber/v2"
)

func FindCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := utils.GetContext()

	defer cancel()

	customerService := services.CustomerService{}

	dto := &pb.CustomerRequest{
		Id: id,
	}

	customer, err := customerService.FindCustomer(ctx, dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
