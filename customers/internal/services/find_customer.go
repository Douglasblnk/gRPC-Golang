package services

import (
	"microservice-poc/customers/internal/models"
	"microservice-poc/customers/internal/repository"
	pbcustomer "microservice-poc/proto"
)

type CustomerService struct {
	pbcustomer.UnimplementedCustomerServiceServer
}

func FindCustomer(id string) (*models.Customer, error) {
	customer, err := repository.FindCustomerByID(id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
