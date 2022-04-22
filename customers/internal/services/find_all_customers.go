package services

import (
	"microservice-poc/customers/internal/models"
	"microservice-poc/customers/internal/repository"
)

func FindAllCustomers() ([]*models.Customer, error) {
	customers, err := repository.FindAllCustomers()

	if err != nil {
		return nil, err
	}

	return customers, nil
}
