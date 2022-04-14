package repository

import (
	"customer/internal/database"
	"customer/internal/models"
)

func CreateCustomer(customer *models.Customer) error {
	if err := database.DB.Select("*").Create(customer).Error; err != nil {
		return err
	}

	return nil
}
