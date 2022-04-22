package repository

import (
	"microservice-poc/customers/internal/database"
	"microservice-poc/customers/internal/exceptions"
	"microservice-poc/customers/internal/models"
)

func FindCustomerByID(id string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := database.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customer, nil
}
