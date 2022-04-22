package repository

import (
	"customers/internal/database"
	"customers/internal/models"
)

func CreateCustomer(customer *models.Customer) error {
	if err := database.DB.Select("*").Create(customer).Error; err != nil {
		return err
	}

	return nil
}
