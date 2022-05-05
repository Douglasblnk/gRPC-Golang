package routes

import (
	"microservice-poc/products/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/:id", handlers.FindProductHandler)
	app.Get("/", handlers.FindAllProductsHandler)
	app.Post("/", handlers.CreateProductHandler)
	app.Put("/:id", handlers.UpdateProductHandler)
	app.Delete("/:id", handlers.DeleteProductHandler)
}
