package services

import (
	"customers/internal/models"
	"customers/internal/repository"
)

func FindAllCustomers() ([]*models.Customer, error) {
	customers, err := repository.FindAllCustomers()

	if err != nil {
		return nil, err
	}

	return customers, nil
}
