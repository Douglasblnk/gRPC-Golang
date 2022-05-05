package restServer

import (
	"fmt"
	"log"
	"microservice-poc/customers/api/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	app := fiber.New()

	routes.SetUpRoutes(app)

	fmt.Println("Listening in port 3000")

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
