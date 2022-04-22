package repository

import (
	"microservice-poc/products/internal/database"
	"microservice-poc/products/internal/models"
)

func CreateProduct(customer *models.Product) error {
	if err := database.DB.Select("*").Create(customer).Error; err != nil {
		return err
	}

	return nil
}
