package main

import (
	"customers/api"
	"customers/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}

	// if err := config.InitgRPC(); err != nil {
	// 	log.Fatal(err)
	// }
}

func main() {
	api.InitServer()
}
