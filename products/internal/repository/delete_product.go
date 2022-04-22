package repository

import (
	"microservice-poc/products/internal/database"
	"microservice-poc/products/internal/models"
)

func DeleteProduct(product *models.Product) error {
	if err := database.DB.Delete(product).Error; err != nil {
		return err
	}

	return nil
}
