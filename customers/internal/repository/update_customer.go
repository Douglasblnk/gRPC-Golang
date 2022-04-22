package repository

import (
	"customers/internal/database"
	"customers/internal/models"
)

func UpdateCustomer(customer *models.Customer) error {
	if err := database.DB.Save(customer).Error; err != nil {
		return err
	}

	return nil
}
