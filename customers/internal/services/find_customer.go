package services

import (
	"context"
	"microservice-poc/customers/internal/repository"
	pb "microservice-poc/customers/proto"
)

type CustomerService struct {
	pb.CustomerServiceServer
}

func (cs *CustomerService) FindCustomer(ctx context.Context, data *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customer, err := repository.FindCustomerByID(data.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.CustomerResponse{
		Id:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}, nil
}
