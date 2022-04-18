package repository

import (
	"products/internal/database"
	"products/internal/exceptions"
	"products/internal/models"
)

func FindAllProducts() ([]*models.Product, error) {
	products := []*models.Product{}

	if err := database.DB.Find(&products).Error; err != nil {
		return nil, exceptions.ErrProductNotFound
	}

	return products, nil
}
