package repository

import (
	"customers/internal/database"
	"customers/internal/exceptions"
	"customers/internal/models"
)

func FindAllCustomers() ([]*models.Customer, error) {
	customers := []*models.Customer{}

	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	return customers, nil
}
