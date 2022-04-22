package repository

import (
	"microservice-poc/customers/internal/database"
	"microservice-poc/customers/internal/models"
)

func DeleteCustomer(customer *models.Customer) error {
	if err := database.DB.Delete(customer).Error; err != nil {
		return err
	}

	return nil
}
