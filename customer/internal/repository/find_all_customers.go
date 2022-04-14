package repository

import (
	"customer/internal/database"
	"customer/internal/exceptions"
	"customer/internal/models"
)

func FindAllCustomers() ([]*models.Customer, error) {
	customers := []*models.Customer{}

	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customers, nil
}
