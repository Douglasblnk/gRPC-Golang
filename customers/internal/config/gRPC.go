package config

import (
	"customers/internal/services"
	"log"
	"net"
	pbcustomer "proto"

	"google.golang.org/grpc"
)

func InitgRPC() error {
	lis, err := net.Listen("tcp", "3001")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	customer := services.CustomerService{}

	grpcServer := grpc.NewServer()

	pbcustomer.RegisterCustomerServiceServer(grpcServer, customer)

	log.Println("Listening on Port: 8200!")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return nil
}
