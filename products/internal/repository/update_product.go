package repository

import (
	"products/internal/database"
	"products/internal/models"
)

func UpdateProduct(product *models.Product) error {
	if err := database.DB.Save(product).Error; err != nil {
		return err
	}

	return nil
}
