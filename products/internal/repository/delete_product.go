package repository

import (
	"products/internal/database"
	"products/internal/models"
)

func DeleteProduct(product *models.Product) error {
	if err := database.DB.Delete(product).Error; err != nil {
		return err
	}

	return nil
}
