package repository

import (
	"customer/internal/database"
	"customer/internal/models"
)

func DeleteCustomer(customer *models.Customer) error {
	if err := database.DB.Delete(customer).Error; err != nil {
		return err
	}

	return nil
}
