package routes

import (
	"microservice-poc/orders/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/:id", handlers.FindOrderHandler)
	app.Get("/", handlers.FindAllOrdersHandler)
	app.Post("/", handlers.CreateOrderHandler)
	// app.Put("/:id", handlers.UpdateProductHandler)
	// app.Delete("/:id", handlers.DeleteProductHandler)
}
