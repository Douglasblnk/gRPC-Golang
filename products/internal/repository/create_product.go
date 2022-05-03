package repository

import (
	"microservice-poc/products/internal/database"
	"microservice-poc/products/internal/models"
)

func CreateProduct(product *models.Product) error {
	if err := database.DB.Select("*").Create(product).Error; err != nil {
		return err
	}

	return nil
}
