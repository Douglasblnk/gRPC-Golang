package services

import (
	"customers/internal/models"
	"customers/internal/repository"
	pbcustomer "proto"
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
