package handlers

import (
	"microservice-poc/customers/api/utils"
	"microservice-poc/products/internal/services"
	pb "microservice-poc/products/proto"

	"github.com/gofiber/fiber/v2"
)

func FindProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := utils.GetContext()

	defer cancel()

	productService := services.ProductService{}

	dto := &pb.ProductRequest{
		Id: id,
	}

	product, err := productService.FindProduct(ctx, dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
