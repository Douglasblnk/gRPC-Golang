package grpcServer

import (
	"fmt"
	"log"
	grpcHandlers "microservice-poc/products/api/grpc/handlers"
	pb "microservice-poc/products/proto"
	"net"

	"google.golang.org/grpc"
)

func InitgRPC() {
	lis, err := net.Listen("tcp", "localhost:3300")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	product := &grpcHandlers.ProductService{}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, product)

	fmt.Println("gRPC server listening on port: 3300")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
