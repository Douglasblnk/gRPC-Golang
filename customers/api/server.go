package api

import (
	"fmt"
	"log"
	"microservice-poc/customers/api/routes"

	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	app := fiber.New()

	routes.SetUpRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening in port 3000")
}
