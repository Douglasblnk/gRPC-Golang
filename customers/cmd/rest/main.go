package main

import (
	"log"
	restServer "microservice-poc/customers/api/rest"
	"microservice-poc/customers/internal/database"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	restServer.InitServer()
}
