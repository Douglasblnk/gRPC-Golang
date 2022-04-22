package repository

import (
	"microservice-poc/customers/internal/database"
	"microservice-poc/customers/internal/exceptions"
	"microservice-poc/customers/internal/models"
)

func FindAllCustomers() ([]*models.Customer, error) {
	customers := []*models.Customer{}

	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customers, nil
}
