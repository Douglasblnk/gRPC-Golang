package main

import (
	"log"
	grpcServer "microservice-poc/products/api/grpc"
	"microservice-poc/products/internal/database"

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
	grpcServer.InitgRPC()
}
