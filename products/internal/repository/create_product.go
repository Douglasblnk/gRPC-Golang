package repository

import (
	"products/internal/database"
	"products/internal/models"
)

func CreateProduct(customer *models.Product) error {
	if err := database.DB.Select("*").Create(customer).Error; err != nil {
		return err
	}

	return nil
}
