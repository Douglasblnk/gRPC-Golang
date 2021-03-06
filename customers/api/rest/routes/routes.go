package routes

import (
	handlers "microservice-poc/customers/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/:id", handlers.FindCustomerHandler)
	app.Get("/", handlers.FindAllCustomersHandler)
	app.Post("/", handlers.CreateCustomerHandler)
	app.Put("/:id", handlers.UpdateCustomerHandler)
	app.Delete("/:id", handlers.DeleteCustomerHandler)
}
