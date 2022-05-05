package main

import (
	"log"
	grpcServer "microservice-poc/customers/api/grpc"
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
	grpcServer.InitgRPC()
}
