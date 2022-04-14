package repository

import (
	"customer/internal/database"
	"customer/internal/models"
)

func UpdateCustomer(customer *models.Customer) error {
	if err := database.DB.Save(customer).Error; err != nil {
		return err
	}

	return nil
}
