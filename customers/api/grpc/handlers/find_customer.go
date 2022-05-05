package grpcHandlers

import (
	"context"
	"microservice-poc/customers/internal/services"
	pb "microservice-poc/customers/proto"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
}

func (cs *CustomerService) FindCustomer(ctx context.Context, data *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customer, err := services.FindCustomer(data.Id)

	if err != nil {
		return nil, err
	}

	return &pb.CustomerResponse{
		Id:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}, nil
}
