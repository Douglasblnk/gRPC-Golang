package repository

import (
	"customer/internal/database"
	"customer/internal/exceptions"
	"customer/internal/models"
)

func FindCustomerByEmail(email string) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := database.DB.Where("email = ?", email).First(customer).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customer, nil
}
