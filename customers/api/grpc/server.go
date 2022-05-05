package grpcServer

import (
	"fmt"
	"log"
	grpcHandlers "microservice-poc/customers/api/grpc/handlers"
	pb "microservice-poc/customers/proto"
	"net"

	"google.golang.org/grpc"
)

func InitgRPC() {
	lis, err := net.Listen("tcp", "localhost:3200")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	customer := &grpcHandlers.CustomerService{}

	grpcServer := grpc.NewServer()

	pb.RegisterCustomerServiceServer(grpcServer, customer)

	fmt.Println("gRPC server listening on port: 3200")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
