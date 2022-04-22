package repository

import (
	"customers/internal/database"
	"customers/internal/exceptions"
	"customers/internal/models"
)

func FindCustomerByID(id string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := database.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customer, nil
}
