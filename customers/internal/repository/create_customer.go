package repository

import (
	"microservice-poc/customers/internal/database"
	"microservice-poc/customers/internal/models"
)

func CreateCustomer(customer *models.Customer) error {
	if err := database.DB.Select("*").Create(customer).Error; err != nil {
		return err
	}

	return nil
}
