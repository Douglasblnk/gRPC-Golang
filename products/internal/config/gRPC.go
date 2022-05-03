package config

import (
	"log"
	"microservice-poc/products/internal/services"
	pb "microservice-poc/products/proto"
	"net"

	"google.golang.org/grpc"
)

func InitgRPC() {
	lis, err := net.Listen("tcp", "localhost:3300")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	product := &services.ProductService{}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, product)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
