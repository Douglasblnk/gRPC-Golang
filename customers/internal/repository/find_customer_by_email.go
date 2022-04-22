package repository

import (
	"customers/internal/database"
	"customers/internal/exceptions"
	"customers/internal/models"
)

func FindCustomerByEmail(email string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := database.DB.Where("email = ?", email).First(customer).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customer, nil
}
