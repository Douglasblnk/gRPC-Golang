package config

import (
	"log"
	"microservice-poc/customers/internal/services"
	pb "microservice-poc/customers/proto"
	"net"

	"google.golang.org/grpc"
)

func InitgRPC() {
	lis, err := net.Listen("tcp", "localhost:3200")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	customer := &services.CustomerService{}

	grpcServer := grpc.NewServer()

	pb.RegisterCustomerServiceServer(grpcServer, customer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
